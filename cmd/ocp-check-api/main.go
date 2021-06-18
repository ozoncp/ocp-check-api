package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-check-api/internal/api"
	"github.com/ozoncp/ocp-check-api/internal/producer"
	prom "github.com/ozoncp/ocp-check-api/internal/prometheus"
	"github.com/ozoncp/ocp-check-api/internal/repo"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	grpczerolog "github.com/philip-bui/grpc-zerolog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	zerolog "github.com/rs/zerolog"
	"google.golang.org/grpc"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

const (
	grpcAddress       = "0.0.0.0:8083"
	prometheusAddress = "0.0.0.0:9100"
)

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

func runGrpcServer(address string) error {
	log := zerolog.New(os.Stdout)
	ctx := context.Background()
	db, err := sqlx.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panic().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	initOpentracing(log)

	repo := repo.NewCheckRepo(db, &log)

	producer, err := producer.NewProducer(ctx)
	if err != nil {
		log.Error().Msgf("failed to create kafka provider: %v", err)
	}

	listen, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	s := grpc.NewServer(grpczerolog.UnaryInterceptorWithLogger(&log))

	prom := prom.NewPrometheus(log)

	http.Handle("/metrics", promhttp.Handler())
	log.Info().Msgf("listen Prometheus on %s", prometheusAddress)
	if err = http.ListenAndServe(prometheusAddress, promhttp.Handler()); err != nil {
		log.Fatal().Err(err).Msgf("failed to listen or serve Prometheus: %v", err)
		return err
	}

	desc.RegisterOcpCheckApiServer(s, api.NewOcpCheckApi(100, log, repo, producer, prom, opentracing.GlobalTracer()))

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
		return err
	}

	return nil
}

func main() {
	_ = runGrpcServer(grpcAddress)
}
