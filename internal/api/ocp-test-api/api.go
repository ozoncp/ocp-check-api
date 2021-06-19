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
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-test-api"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	batchSize uint
	log       zerolog.Logger
	repo      repo.TestRepo
	producer  producer.Producer
	prom      prometheus.Prometheus
	tracer    opentracing.Tracer
	desc.UnimplementedOcpTestApiServer
}

func (a *api) SendEvent(event producer.TestEvent) error {
	return a.producer.SendTestEvent(event)
}

func (a *api) ListTests(ctx context.Context,
	req *desc.ListTestsRequest,
) (*desc.ListTestsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var Tests []models.Test
	var err error

	Tests, err = a.repo.ListTests(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var pbTests []*desc.Test
	for _, Test := range Tests {
		pbTests = append(pbTests, &desc.Test{
			Id:     Test.ID,
			TaskID: Test.TaskID,
			Input:  Test.Input,
			Output: Test.Output,
		})
	}
	return &desc.ListTestsResponse{Tests: pbTests}, err
}

func (a *api) DescribeTest(
	ctx context.Context,
	req *desc.DescribeTestRequest,
) (*desc.DescribeTestResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var Test *models.Test
	var err error

	Test, err = a.repo.DescribeTest(ctx, req.TestId)
	if err != nil {
		switch {
		case err == repo.TestNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Unknown, err.Error())
		}
	}

	pbTest := &desc.Test{
		Id:     Test.ID,
		TaskID: Test.TaskID,
		Input:  Test.Input,
		Output: Test.Output,
	}

	return &desc.DescribeTestResponse{Test: pbTest}, nil
}

func (a *api) CreateTest(ctx context.Context,
	req *desc.CreateTestRequest,
) (*desc.CreateTestResponse, error) {
	var err error
	if err = req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	Test := models.Test{
		TaskID: req.TaskID,
		Input:  req.Input,
		Output: req.Output,
	}

	var id uint64
	if id, err = a.repo.CreateTest(ctx, Test); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if id != 0 {
		Test.ID = id
		_ = a.producer.SendTestEvent(producer.TestEvent{Type: producer.Created, Event: Test})
		a.prom.IncCreateTest("success")
	}

	a.log.Info().Msgf("New Test created: id=%v", id)

	return &desc.CreateTestResponse{TestId: id}, nil
}

func (a *api) MultiCreateTest(ctx context.Context,
	req *desc.MultiCreateTestRequest,
) (*desc.MultiCreateTestResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, a.tracer, "MultiCreateTest parent")
	defer span.Finish()

	var err error
	if err = req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	Tests := make([]models.Test, 0, len(req.Tests))
	for _, Test := range req.Tests {
		newTest := models.Test{
			TaskID: Test.TaskID,
			Input:  Test.Input,
			Output: Test.Output,
		}
		Tests = append(Tests, newTest)
	}

	batches, err := utils.SplitTestsToBulks(Tests, a.batchSize)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	var totalCreatedTests = uint64(0)
	for _, batch := range batches {
		childSpan, _ := opentracing.StartSpanFromContextWithTracer(ctx, a.tracer, "MultiCreateTest batch")
		childSpan.SetTag("batchSize", fmt.Sprintf("%v", len(batch)))
		defer childSpan.Finish()

		createdTests, err := a.repo.MultiCreateTest(ctx, batch)
		if err != nil {
			return &desc.MultiCreateTestResponse{Created: totalCreatedTests}, status.Error(codes.Unknown, err.Error())
		}

		if len(batch) == len(createdTests) {
			for idx, TestId := range createdTests {
				batch[idx].ID = TestId
				_ = a.producer.SendTestEvent(producer.TestEvent{Type: producer.Created, Event: batch[idx]})
				a.prom.IncCreateTest("success")
			}
		}
		totalCreatedTests += uint64(len(createdTests))
	}

	return &desc.MultiCreateTestResponse{Created: totalCreatedTests}, nil
}

func (a *api) UpdateTest(ctx context.Context,
	req *desc.UpdateTestRequest,
) (*desc.UpdateTestResponse, error) {

	updatedTest := models.Test{
		ID:     req.Test.Id,
		TaskID: req.Test.TaskID,
		Input:  req.Test.Input,
		Output: req.Test.Output,
	}

	updated, err := a.repo.UpdateTest(ctx, updatedTest)
	switch {
	case err == repo.TestNotFound:
		return nil, status.Error(codes.NotFound, err.Error())
	case err != nil:
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if updated {
		_ = a.producer.SendTestEvent(producer.TestEvent{Type: producer.Updated, Event: updatedTest})
		a.prom.IncUpdateTest("success")
	}

	return &desc.UpdateTestResponse{
		Updated: updated,
	}, nil
}

func (a *api) RemoveTest(ctx context.Context,
	req *desc.RemoveTestRequest,
) (*desc.RemoveTestResponse, error) {

	var found = true

	err := a.repo.RemoveTest(ctx, req.TestId)
	switch {
	case err == repo.TestNotFound:
		found = false
	case err != nil:
		return nil, status.Error(codes.Unknown, err.Error())
	}

	if found {
		deletedTest := models.Test{ID: req.TestId}
		_ = a.producer.SendTestEvent(producer.TestEvent{Type: producer.Deleted, Event: deletedTest})
		a.prom.IncDeleteTest("success")
	}

	return &desc.RemoveTestResponse{
		Deleted: found,
	}, nil
}

func NewOcpTestApi(batchSize uint, log zerolog.Logger, repo repo.TestRepo, producer producer.Producer, prom prometheus.Prometheus, tracer opentracing.Tracer) desc.OcpTestApiServer {
	return &api{
		batchSize: batchSize,
		log:       log,
		repo:      repo,
		producer:  producer,
		prom:      prom,
		tracer:    tracer}
}
