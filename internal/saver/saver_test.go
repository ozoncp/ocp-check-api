package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-check-api/internal/mocks"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		err error

		ctrl *gomock.Controller

		mockFlusher *mocks.MockCheckFlusher

		check models.Check
		s     saver.Saver
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockFlusher = mocks.NewMockCheckFlusher(ctrl)

		s = saver.NewSaver(1000, 2, mockFlusher)
	})

	JustBeforeEach(func() {
		err = s.Save(check)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("operation canceled", func() {

		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Return(nil).Times(1)
		})

		JustBeforeEach(func() {
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("alarm is occurring", func() {

		BeforeEach(func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Return(nil).MinTimes(1).MaxTimes(2)
		})

		JustBeforeEach(func() {
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})
})
