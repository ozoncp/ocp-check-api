package api_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"

	"github.com/ozoncp/ocp-check-api/internal/api"
	"github.com/ozoncp/ocp-check-api/internal/mocks"
	"github.com/ozoncp/ocp-check-api/internal/producer"
	"github.com/ozoncp/ocp-check-api/internal/prometheus"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
)

var _ = Describe("Api", func() {
	var (
		ctrl     *gomock.Controller
		ctx      context.Context
		log      zerolog.Logger
		producer producer.Producer
		metrics  prometheus.Prometheus

		mockRepo *mocks.MockCheckRepo
		grpcApi  desc.OcpCheckApiServer

		createReq *desc.CreateCheckRequest
		err       error
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockRepo = mocks.NewMockCheckRepo(ctrl)
		grpcApi = api.NewOcpCheckApi(100, log, mockRepo, producer, metrics)

		createReq = &desc.CreateCheckRequest{SolutionID: 2, TestID: 3, RunnerID: 4, Success: false}
	})

	Context("insert single check into database", func() {
		const (
			newId = uint64(100)
		)

		var (
			createResp *desc.CreateCheckResponse
		)
		BeforeEach(func() {
			mockRepo.EXPECT().AddCheck(gomock.Any()).MinTimes(1).Return(newId, nil)
			createResp, err = grpcApi.CreateCheck(ctx, createReq)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResp.CheckId).Should(BeNumerically(">", 0))
			Expect(createResp.CheckId).Should(Equal(newId))
		})
	})
})
