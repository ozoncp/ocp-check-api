package alarmer

import (
	"context"
	"time"
)

type Alarmer interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

type alarmer struct {
	ctx    context.Context
	period time.Duration
	alarms chan struct{}
	done   chan struct{}
}

func NewAlarmer(ctx context.Context, period time.Duration) Alarmer {

	if period <= 0 {
		return nil
	}

	alarms := make(chan struct{})
	done := make(chan struct{})

	return &alarmer{
		ctx:    ctx,
		period: period,
		alarms: alarms,
		done:   done,
	}
}

func (a *alarmer) Init() {
	go func() {
		timer := time.After(a.period)
		for {
			select {
			case <-timer:
				a.alarms <- struct{}{}
				timer = time.After(a.period)
			case <-a.ctx.Done():
				close(a.alarms)
				a.done <- struct{}{}
				return
			}
		}
	}()
}

func (a *alarmer) Alarm() <-chan struct{} {
	return a.alarms
}

func (a *alarmer) Close() {
	<-a.done
}
