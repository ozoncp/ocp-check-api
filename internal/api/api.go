package api

import (
	"context"

	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	errCheckNotFound     = "check not found"
	errNoChecksAvailable = "no checks"
	errInvalidArg        = "invalid argument"
)

type api struct {
	desc.UnimplementedOcpCheckApiServer
}

func (a *api) ListChecks(ctx context.Context,
	req *desc.ListChecksRequest,
) (*desc.ListChecksResponse, error) {
	err := status.Error(codes.NotFound, errNoChecksAvailable)
	return nil, err
}

func (a *api) DescribeCheck(
	ctx context.Context,
	req *desc.DescribeCheckRequest,
) (*desc.DescribeCheckResponse, error) {
	err := status.Error(codes.NotFound, errCheckNotFound)
	return nil, err
}

func (a *api) CreateCheck(ctx context.Context,
	req *desc.CreateCheckRequest,
) (*desc.CreateCheckResponse, error) {
	err := status.Error(codes.InvalidArgument, errInvalidArg)
	return nil, err
}

func (a *api) RemoveCheck(ctx context.Context,
	req *desc.RemoveCheckRequest,
) (*desc.RemoveCheckResponse, error) {
	err := status.Error(codes.NotFound, errCheckNotFound)
	return nil, err
}

func NewOcpCheckApi() desc.OcpCheckApiServer {
	return &api{}
}
