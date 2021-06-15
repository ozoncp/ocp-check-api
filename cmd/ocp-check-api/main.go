package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	api "github.com/ozoncp/ocp-check-api/internal/api"
	"github.com/ozoncp/ocp-check-api/internal/repo"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	grpczerolog "github.com/philip-bui/grpc-zerolog"
	zerolog "github.com/rs/zerolog"
	"google.golang.org/grpc"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

const (
	grpcAddress = "0.0.0.0:8083"
)

func runGrpcServer(address string) error {
	log := zerolog.New(os.Stdout)
	ctx := context.Background()
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panic().Err(err).Msg("Unable to connect to database")
	}
	defer db.Close()

	repo := repo.NewCheckRepo(&ctx, db, &log)

	listen, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	s := grpc.NewServer(grpczerolog.UnaryInterceptorWithLogger(&log))

	desc.RegisterOcpCheckApiServer(s, api.NewOcpCheckApi(repo))

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}

	return nil
}

func main() {
	_ = runGrpcServer(grpcAddress)
}
