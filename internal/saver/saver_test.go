package saver_test

import (
	"context"

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
		ctx  context.Context

		mockFlusher *mocks.MockCheckFlusher
		mockAlarmer *mocks.MockAlarmer

		check  models.Check
		s      saver.Saver
		alarms chan struct{}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockAlarmer = mocks.NewMockAlarmer(ctrl)
		mockFlusher = mocks.NewMockCheckFlusher(ctrl)

		alarms = make(chan struct{})
		mockAlarmer.EXPECT().Alarm().Return(alarms).AnyTimes()

		s = saver.NewSaver(uint(1000), mockAlarmer, mockFlusher)
	})

	JustBeforeEach(func() {
		s.Init(ctx)
		err = s.Save(ctx, check)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("operation canceled", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil)
		})

		JustBeforeEach(func() {
			cancelFunc()
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("alarm is occurring", func() {
		var (
			cancelFunc context.CancelFunc
		)
		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1).MaxTimes(2)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
			cancelFunc()
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})
})
