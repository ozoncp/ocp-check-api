package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-check-api/internal/api/ocp-check-api"
	apit "github.com/ozoncp/ocp-check-api/internal/api/ocp-test-api"
	"github.com/ozoncp/ocp-check-api/internal/producer"
	prom "github.com/ozoncp/ocp-check-api/internal/prometheus"
	repo "github.com/ozoncp/ocp-check-api/internal/repo"
	descc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	desct "github.com/ozoncp/ocp-check-api/pkg/ocp-test-api"
	grpczerolog "github.com/philip-bui/grpc-zerolog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	zerolog "github.com/rs/zerolog"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

var (
	InvalidConfigError = errors.New("Unable to create logger: console and file loggers are disabled. Check config file...")
	GitCommit          string
	ProtocolRevision   = "V1"
	BuildDateTime      string
)

type config struct {
	Grpc struct {
		Address string `yaml:"address"`
	}
	Prometheus struct {
		Address string `yaml:"address"`
	}
	Kafka struct {
		Brokers []string `yaml:"brokers"`
	}
	Database struct {
		Url string `yaml:"url"`
	}
	Log struct {
		Console struct {
			Enable bool `yaml:"enable"`
		}
		File struct {
			Enable bool   `yaml:"enable"`
			Path   string `yaml:"path"`
		}
	}
}

var configDefaults = map[string]interface{}{
	"grpc.address":       "0.0.0.0:8083",
	"prometheus.address": "0.0.0.0:9100",
	"kafka.brokers":      "127.0.0.1:9092",
	"database.url":       "postgres://postgres:postgres@localhost:5432/postgres",
}

// Read the config file from the current directory and marshal it into the config struct.
func getConfig() *config {
	viper.AddConfigPath("./")
	viper.SetConfigFile("config.yml")

	for k, v := range configDefaults {
		viper.SetDefault(k, v)
	}

	viper.SetEnvPrefix("ocp_check_api")

	bindEnvs := []string{"grpc_address", "prometheus_address", "kafka_brokers", "database_url"}

	for _, env := range bindEnvs {
		_ = viper.BindEnv(env)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return conf
}

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func initOpentracing(log zerolog.Logger) {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: "ocp-check-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		log.Error().Err(err).Msg("")
	}

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
}

func runGrpcServer(cfg *config) error {
	if !cfg.Log.Console.Enable && !cfg.Log.File.Enable {
		fmt.Printf("%v", InvalidConfigError.Error())
		return InvalidConfigError
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	var log zerolog.Logger

	if cfg.Log.File.Enable {
		logFile, err := os.OpenFile(cfg.Log.File.Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}

		if cfg.Log.Console.Enable {
			multi := io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}, logFile)
			log = zerolog.New(multi).With().Timestamp().Logger()
		} else {
			log = zerolog.New(logFile).With().Timestamp().Logger()
		}
	}

	ctx, done := context.WithCancel(context.Background())
	g, gctx := errgroup.WithContext(ctx)

	db, err := sqlx.Open("pgx", cfg.Database.Url)
	if err != nil {
		log.Panic().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	initOpentracing(log)

	checkRepo := repo.NewCheckRepo(db, &log)
	testRepo := repo.NewTestRepo(db, &log)

	producer, err := producer.NewProducer(ctx, cfg.Kafka.Brokers)
	if err != nil {
		log.Error().Msgf("failed to create kafka provider: %v", err)
	}

	listen, err := net.Listen("tcp4", cfg.Grpc.Address)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	prom := prom.NewPrometheus(log)

	metricServer := &http.Server{Addr: cfg.Prometheus.Address}
	http.Handle("/metrics", promhttp.Handler())
	s := grpc.NewServer(grpczerolog.UnaryInterceptorWithLogger(&log))

	go func() {
		log.Info().Msgf("listen Prometheus on %s", cfg.Prometheus.Address)
		err = metricServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msgf("failed to listen or serve Prometheus: %v", err)
		}
	}()

	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

		select {
		case sig := <-signalChannel:
			fmt.Printf("Close by ctrl+c: %s\n", sig)
			fmt.Printf("shutting down Prometheus\n")
			if err = metricServer.Shutdown(ctx); err != nil {
				log.Error().Err(err).Msgf("error during shutdown")
			}
			s.GracefulStop()
			done()
		case <-gctx.Done():
			return gctx.Err()
		}

		return nil
	})

	g.Go(func() error {
		buildInfo := api.BuildInfo{
			GitCommit:        GitCommit,
			ProtocolRevision: ProtocolRevision,
			BuildDateTime:    BuildDateTime,
		}

		descc.RegisterOcpCheckApiServer(s, api.NewOcpCheckApi(buildInfo, 100, log, checkRepo, producer, prom, opentracing.GlobalTracer()))
		desct.RegisterOcpTestApiServer(s, apit.NewOcpTestApi(100, log, testRepo, producer, prom, opentracing.GlobalTracer()))

		if err := s.Serve(listen); err != nil {
			log.Fatal().Err(err).Msg("failed to serve")
			return err
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatal().Err(err).Msg("failed to wait goroutine group")
	}

	log.Info().Msg("graceful shutdown successfully finished")

	return nil
}

func main() {
	config := getConfig()
	_ = runGrpcServer(config)
}
