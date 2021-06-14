package main

import (
	"fmt"
	"net"
	"os"

	api "github.com/ozoncp/ocp-check-api/internal/api"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	grpczerolog "github.com/philip-bui/grpc-zerolog"
	zerolog "github.com/rs/zerolog"
	"google.golang.org/grpc"
)

const (
	grpcAddress = "0.0.0.0:8083"
)

func Greeting(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func runGrpcServer(address string) error {
	// console Zerolog instance.
	log := zerolog.New(os.Stdout)

	listen, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	s := grpc.NewServer(grpczerolog.UnaryInterceptorWithLogger(&log))

	desc.RegisterOcpCheckApiServer(s, api.NewOcpCheckApi())

	if err := s.Serve(listen); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}

	return nil
}

func main() {
	_ = runGrpcServer(grpcAddress)
}
