package api

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/producer"
	"github.com/ozoncp/ocp-check-api/internal/prometheus"
	"github.com/ozoncp/ocp-check-api/internal/repo"
	"github.com/ozoncp/ocp-check-api/internal/utils"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	batchSize uint
	log       zerolog.Logger
	repo      repo.CheckRepo
	producer  producer.Producer
	prom      prometheus.Prometheus
	tracer    opentracing.Tracer
	desc.UnimplementedOcpCheckApiServer
}

func (a *api) SendEvent(event producer.CheckEvent) error {
	return a.producer.SendEvent(event)
}

func (a *api) ListChecks(ctx context.Context,
	req *desc.ListChecksRequest,
) (*desc.ListChecksResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var checks []models.Check
	var err error

	checks, err = a.repo.ListChecks(ctx, req.Limit, req.Offset)
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

	check, err = a.repo.DescribeCheck(ctx, req.CheckId)
	if err != nil {
		switch {
		case err == repo.CheckNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
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
	if id, err = a.repo.CreateCheck(ctx, check); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if id != 0 {
		check.ID = id
		_ = a.producer.SendEvent(producer.CheckEvent{Type: producer.Created, Event: check})
		a.prom.IncCreateCheck("success")
	}

	a.log.Info().Msgf("New check created: id=%v", id)

	return &desc.CreateCheckResponse{CheckId: id}, nil
}

func (a *api) MultiCreateCheck(ctx context.Context,
	req *desc.MultiCreateCheckRequest,
) (*desc.MultiCreateCheckResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, a.tracer, "MultiCreateCheck parent")
	defer span.Finish()

	var err error
	if err = req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	checks := make([]models.Check, 0, len(req.Checks))
	for _, check := range req.Checks {
		newCheck := models.Check{
			SolutionID: check.SolutionID,
			RunnerID:   check.RunnerID,
			TestID:     check.TestID,
			Success:    check.Success,
		}
		checks = append(checks, newCheck)
	}

	batches, err := utils.SplitChecksToBulks(checks, a.batchSize)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var totalCreatedChecks = uint64(0)
	for _, batch := range batches {
		childSpan, _ := opentracing.StartSpanFromContextWithTracer(ctx, a.tracer, "MultiCreateCheck batch")
		childSpan.SetTag("batchSize", fmt.Sprintf("%v", len(batch)))
		defer childSpan.Finish()

		createdChecks, err := a.repo.MultiCreateCheck(ctx, batch)
		if err != nil {
			return &desc.MultiCreateCheckResponse{Created: totalCreatedChecks}, status.Error(codes.Unknown, err.Error())
		}
		totalCreatedChecks += createdChecks
	}

	return &desc.MultiCreateCheckResponse{Created: totalCreatedChecks}, nil
}

func (a *api) UpdateCheck(ctx context.Context,
	req *desc.UpdateCheckRequest,
) (*desc.UpdateCheckResponse, error) {

	updatedCheck := models.Check{
		ID:         req.Check.Id,
		SolutionID: req.Check.SolutionID,
		RunnerID:   req.Check.RunnerID,
		TestID:     req.Check.TestID,
		Success:    req.Check.Success,
	}

	updated, err := a.repo.UpdateCheck(ctx, updatedCheck)
	switch {
	case err == repo.CheckNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case err != nil:
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if updated {
		_ = a.producer.SendEvent(producer.CheckEvent{Type: producer.Updated, Event: updatedCheck})
		a.prom.IncUpdateCheck("success")
	}

	return &desc.UpdateCheckResponse{
		Updated: updated,
	}, nil
}

func (a *api) RemoveCheck(ctx context.Context,
	req *desc.RemoveCheckRequest,
) (*desc.RemoveCheckResponse, error) {

	var found = true

	err := a.repo.RemoveCheck(ctx, req.CheckId)
	switch {
	case err == repo.CheckNotFound:
		found = false
	case err != nil:
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if found {
		deletedCheck := models.Check{ID: req.CheckId}
		_ = a.producer.SendEvent(producer.CheckEvent{Type: producer.Deleted, Event: deletedCheck})
		a.prom.IncDeleteCheck("success")
	}

	return &desc.RemoveCheckResponse{
		Deleted: found,
	}, nil
}

func NewOcpCheckApi(batchSize uint, log zerolog.Logger, repo repo.CheckRepo, producer producer.Producer, prom prometheus.Prometheus, tracer opentracing.Tracer) desc.OcpCheckApiServer {
	return &api{
		batchSize: batchSize,
		log:       log,
		repo:      repo,
		producer:  producer,
		prom:      prom,
		tracer:    tracer}
}
