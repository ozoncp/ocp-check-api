package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/mocks"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		timeOutError = errors.New("timeout elapsed")
		ctrl         *gomock.Controller
		ctx          context.Context

		mockRepo  *mocks.MockTestRepo
		f         flusher.TestFlusher
		chunkSize = 2
		tests     []models.Test
		remained  []models.Test
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockRepo = mocks.NewMockTestRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewTestFlusher(chunkSize, mockRepo)
		remained = f.Flush(ctx, tests)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("save all tests in repo", func() {
		BeforeEach(func() {
			tests = []models.Test{{}, {}, {}}
			gomock.InOrder(
				mockRepo.EXPECT().MultiCreateTest(ctx, gomock.Any()).Return(uint64(chunkSize), nil),
				mockRepo.EXPECT().MultiCreateTest(ctx, gomock.Any()).Return(uint64(len(tests)-chunkSize), nil),
			)
		})

		It("", func() {
			Expect(remained).Should(BeNil())
		})
	})

	Context("error of saving all tests in repo", func() {
		BeforeEach(func() {
			tests = []models.Test{{}, {}}

			mockRepo.EXPECT().MultiCreateTest(ctx, gomock.Len(chunkSize)).Return(uint64(0), timeOutError)
		})

		It("", func() {
			Expect(remained).ShouldNot(BeNil())
			Expect(len(remained)).To(Equal(len(tests)))
		})
	})

	Context("save tests in repo partially", func() {
		BeforeEach(func() {
			tests = []models.Test{{}, {}, {}, {}, {}}

			gomock.InOrder(
				mockRepo.EXPECT().MultiCreateTest(ctx, gomock.Len(chunkSize)).Return(uint64(chunkSize), nil),
				mockRepo.EXPECT().MultiCreateTest(ctx, gomock.Len(chunkSize)).Return(uint64(0), timeOutError),
			)
		})

		It("", func() {
			Expect(remained).ShouldNot(BeNil())
			Expect(len(remained)).To(Equal(len(tests) - chunkSize))
		})
	})
})
