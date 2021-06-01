package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/mocks"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		ctrl *gomock.Controller

		mockRepo  *mocks.MockTestRepo
		err       error
		f         flusher.TestFlusher
		chunkSize = 2
		tests     []models.Test
		rest      []models.Test
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockTestRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewTestFlusher(chunkSize, mockRepo)
		rest = f.Flush(tests)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all tasks", func() {
		BeforeEach(func() {
			tests = []models.Test{{}, {}, {}}

			mockRepo.EXPECT().AddTests(gomock.Any()).Return(nil).MinTimes(2)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo save half of tasks", func() {
		BeforeEach(func() {
			tests = []models.Test{{}, {}, {}}
			mockRepo.EXPECT().AddTests(gomock.Any()).Return().MinTimes(1)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})
})
