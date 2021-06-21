package api_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/rs/zerolog"
	"github.com/rzajac/zltest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/ozoncp/ocp-check-api/internal/api/ocp-check-api"
	"github.com/ozoncp/ocp-check-api/internal/mocks"
	repo "github.com/ozoncp/ocp-check-api/internal/repo"
	desc "github.com/ozoncp/ocp-check-api/pkg/ocp-check-api"
)

var _ = Describe("Api", func() {
	var (
		ctrl *gomock.Controller
		ctx  context.Context
		log  zerolog.Logger

		mockProducer   *mocks.MockProducer
		mockPrometheus *mocks.MockPrometheus
		mockRepo       *mocks.MockCheckRepo
		mockTracer     *mocktracer.MockTracer

		grpcApi desc.OcpCheckApiServer

		createReq      *desc.CreateCheckRequest
		multiCreateReq *desc.MultiCreateCheckRequest
		updateReq      *desc.UpdateCheckRequest
		removeReq      *desc.RemoveCheckRequest

		err       error
		someError error

		batchSize   = uint(100)
		mockZeroLog *zltest.Tester
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockProducer = mocks.NewMockProducer(ctrl)
		mockPrometheus = mocks.NewMockPrometheus(ctrl)
		mockRepo = mocks.NewMockCheckRepo(ctrl)
		mockTracer = mocktracer.New()

		mockZeroLog = zltest.New(GinkgoT())
		log = zerolog.New(mockZeroLog).With().Timestamp().Logger()

		createReq = &desc.CreateCheckRequest{SolutionID: 2, TestID: 3, RunnerID: 4, Success: false}

		checks := make([]*desc.CreateCheckRequest, 0)
		checks = append(checks, createReq)
		checks = append(checks, createReq)
		multiCreateReq = &desc.MultiCreateCheckRequest{Checks: checks}

		updateReq = &desc.UpdateCheckRequest{Check: &desc.Check{Id: 1}}

		removeReq = &desc.RemoveCheckRequest{CheckId: updateReq.Check.Id}

		someError = errors.New("some error")

		grpcApi = api.NewOcpCheckApi(api.BuildInfo{}, batchSize, log, mockRepo, mockProducer, mockPrometheus, mockTracer)
	})

	Context("create check: success", func() {
		const (
			newId = uint64(100)
		)

		var (
			createResp *desc.CreateCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(1)
			mockPrometheus.EXPECT().IncCreateCheck(gomock.Any()).Times(1)
			mockRepo.EXPECT().CreateCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(newId, nil)
			createResp, err = grpcApi.CreateCheck(ctx, createReq)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResp.CheckId).Should(BeNumerically(">", 0))
			Expect(createResp.CheckId).Should(Equal(newId))
			entry := mockZeroLog.LastEntry()
			Expect(entry).ShouldNot(BeNil())
			entry.ExpLevel(zerolog.InfoLevel)
			entry.ExpMsg(fmt.Sprintf("New check created: id=%v", newId))
		})
	})

	Context("create check: failure", func() {
		var (
			createResp *desc.CreateCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(0)
			mockPrometheus.EXPECT().IncCreateCheck(gomock.Any()).Times(0)
			mockRepo.EXPECT().CreateCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(uint64(0), someError)
			createResp, err = grpcApi.CreateCheck(ctx, createReq)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResp).Should(BeNil())
			Expect(err).Should(Equal(status.Error(codes.Unknown, someError.Error())))
			entry := mockZeroLog.LastEntry()
			Expect(entry).Should(BeNil())
		})
	})

	Context("multicreate check: success", func() {
		var (
			multiCreateResp *desc.MultiCreateCheckResponse
			created         []uint64
		)

		BeforeEach(func() {
			created = make([]uint64, len(multiCreateReq.Checks))
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(len(multiCreateReq.Checks))
			mockPrometheus.EXPECT().IncCreateCheck(gomock.Any()).Times(len(multiCreateReq.Checks))
			mockRepo.EXPECT().MultiCreateCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(created, nil)
			multiCreateResp, err = grpcApi.MultiCreateCheck(ctx, multiCreateReq)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResp.Created).Should(Equal(uint64(len(created))))
		})
	})

	Context("multicreate check: failure", func() {
		var (
			multiCreateResp *desc.MultiCreateCheckResponse
			created         []uint64
		)

		BeforeEach(func() {
			created = []uint64{}
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(1)
			mockPrometheus.EXPECT().IncCreateCheck(gomock.Any()).Times(1)
			mockRepo.EXPECT().MultiCreateCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(created, someError)
			multiCreateResp, err = grpcApi.MultiCreateCheck(ctx, multiCreateReq)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResp).ShouldNot(BeNil())
			Expect(multiCreateResp.Created).Should(Equal(uint64(len(created))))
			Expect(err).Should(Equal(status.Error(codes.Unknown, someError.Error())))
		})
	})

	Context("multicreate check: failure after some successful operations", func() {
		var (
			multiCreateResp *desc.MultiCreateCheckResponse
			batchSize       = uint(1)
		)

		BeforeEach(func() {
			grpcApi = api.NewOcpCheckApi(api.BuildInfo{}, batchSize, log, mockRepo, mockProducer, mockPrometheus, mockTracer)

			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(1)
			mockPrometheus.EXPECT().IncCreateCheck(gomock.Any()).Times(1)
			gomock.InOrder(
				mockRepo.EXPECT().MultiCreateCheck(gomock.Any(), gomock.Any()).Return([]uint64{uint64(batchSize)}, nil),
				mockRepo.EXPECT().MultiCreateCheck(gomock.Any(), gomock.Any()).Return([]uint64{}, someError),
			)
			multiCreateResp, err = grpcApi.MultiCreateCheck(ctx, multiCreateReq)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResp).ShouldNot(BeNil())
			Expect(err).Should(Equal(status.Error(codes.Unknown, someError.Error())))
			Expect(multiCreateResp.Created).Should(Equal(uint64(batchSize)))
		})
	})

	Context("update check: success", func() {
		var (
			updateResp *desc.UpdateCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(1)
			mockPrometheus.EXPECT().IncUpdateCheck(gomock.Any()).Times(1)
			mockRepo.EXPECT().UpdateCheck(gomock.Any(), gomock.Any()).Return(true, nil)
			updateResp, err = grpcApi.UpdateCheck(ctx, updateReq)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(updateResp).ShouldNot(BeNil())
			Expect(updateResp.Updated).Should(Equal(true))
		})
	})

	Context("update check: not found", func() {
		var (
			updateResp *desc.UpdateCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(0)
			mockPrometheus.EXPECT().IncUpdateCheck(gomock.Any()).Times(0)
			mockRepo.EXPECT().UpdateCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(false, repo.ErrCheckNotFound)
			updateResp, err = grpcApi.UpdateCheck(ctx, updateReq)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(updateResp).Should(BeNil())
			Expect(err).Should(Equal(status.Error(codes.NotFound, repo.ErrCheckNotFound.Error())))
			entry := mockZeroLog.LastEntry()
			Expect(entry).Should(BeNil())
		})
	})

	Context("update check: failure", func() {
		var (
			updateResp *desc.UpdateCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(0)
			mockPrometheus.EXPECT().IncUpdateCheck(gomock.Any()).Times(0)
			mockRepo.EXPECT().UpdateCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(false, someError)
			updateResp, err = grpcApi.UpdateCheck(ctx, updateReq)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(updateResp).Should(BeNil())
			Expect(err).Should(Equal(status.Error(codes.Unknown, someError.Error())))
			entry := mockZeroLog.LastEntry()
			Expect(entry).Should(BeNil())
		})
	})

	Context("remove check: success", func() {
		var (
			removeResp *desc.RemoveCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(1)
			mockPrometheus.EXPECT().IncDeleteCheck(gomock.Any()).Times(1)
			mockRepo.EXPECT().RemoveCheck(gomock.Any(), gomock.Any()).Return(nil)
			removeResp, err = grpcApi.RemoveCheck(ctx, removeReq)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResp).ShouldNot(BeNil())
			Expect(removeResp.Deleted).Should(Equal(true))
		})
	})

	Context("remove check: not found", func() {
		var (
			removeResp *desc.RemoveCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(0)
			mockPrometheus.EXPECT().IncDeleteCheck(gomock.Any()).Times(0)
			mockRepo.EXPECT().RemoveCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(repo.ErrCheckNotFound)
			removeResp, err = grpcApi.RemoveCheck(ctx, removeReq)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResp).ShouldNot(BeNil())
			Expect(removeResp.Deleted).Should(Equal(false))
			entry := mockZeroLog.LastEntry()
			Expect(entry).Should(BeNil())
		})
	})

	Context("remove check: failure", func() {
		var (
			removeResp *desc.RemoveCheckResponse
		)

		BeforeEach(func() {
			mockProducer.EXPECT().SendCheckEvent(gomock.Any()).Times(0)
			mockPrometheus.EXPECT().IncDeleteCheck(gomock.Any()).Times(0)
			mockRepo.EXPECT().RemoveCheck(gomock.Any(), gomock.Any()).MinTimes(1).Return(someError)
			removeResp, err = grpcApi.RemoveCheck(ctx, removeReq)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(removeResp).Should(BeNil())
			Expect(err).Should(Equal(status.Error(codes.Unknown, someError.Error())))
			entry := mockZeroLog.LastEntry()
			Expect(entry).Should(BeNil())
		})
	})
})
