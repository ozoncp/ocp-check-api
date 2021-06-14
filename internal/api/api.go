package api

import (
	"context"

	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/repo"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	repo *repo.CheckRepo
	desc.UnimplementedOcpCheckApiServer
}

func (a *api) ListChecks(ctx context.Context,
	req *desc.ListChecksRequest,
) (*desc.ListChecksResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var checks []models.Check
	var err error
	checks, err = (*a.repo).ListChecks(req.Limit, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var pbChecks []*desc.Check
	for _, check := range checks {
		pbChecks = append(pbChecks, &desc.Check{
			Id:         check.ID,
			SolutionID: check.SolutionID,
			TestID:     check.TestID,
			RunnerID:   check.RunnerID,
			Success:    check.Success,
		})
	}
	return &desc.ListChecksResponse{Checks: pbChecks}, err
}

func (a *api) DescribeCheck(
	ctx context.Context,
	req *desc.DescribeCheckRequest,
) (*desc.DescribeCheckResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var check *models.Check
	var err error

	check, err = (*a.repo).DescribeCheck(req.CheckId)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	pbCheck := &desc.Check{
		Id:         check.ID,
		SolutionID: check.SolutionID,
		TestID:     check.TestID,
		RunnerID:   check.RunnerID,
		Success:    check.Success,
	}
	return &desc.DescribeCheckResponse{Check: pbCheck}, nil
}

func (a *api) CreateCheck(ctx context.Context,
	req *desc.CreateCheckRequest,
) (*desc.CreateCheckResponse, error) {
	var err error
	if err = req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	check := models.Check{
		SolutionID: req.SolutionID,
		RunnerID:   req.RunnerID,
		TestID:     req.TestID,
		Success:    req.Success,
	}

	var id uint64
	if id, err = (*a.repo).AddCheck(check); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &desc.CreateCheckResponse{CheckId: id}, nil
}

func (a *api) RemoveCheck(ctx context.Context,
	req *desc.RemoveCheckRequest,
) (*desc.RemoveCheckResponse, error) {
	if err := (*a.repo).RemoveCheck(req.CheckId); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return &desc.RemoveCheckResponse{
		Found: true,
	}, nil
}

func NewOcpCheckApi(repo *repo.CheckRepo) desc.OcpCheckApiServer {
	return &api{repo: repo}
}
