package saver

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-check-api/internal/alarmer"
	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, check models.Check) error
	Close()
}

// Saver с поддержкой периодического сохранения
type checkSaverPeriodic struct {
	alarmer alarmer.Alarmer
	flusher flusher.CheckFlusher
	checks  chan models.Check
	done    chan struct{}
}

func (c *checkSaverPeriodic) Init(ctx context.Context) {
	var checks []models.Check

	// Flush на закрытие или по тику
	go func(ctx context.Context) {
		for {
			select {
			case e := <-c.checks:
				checks = append(checks, e)

			case <-ctx.Done():
				_ = (c.flusher).Flush(ctx, checks)
				c.done <- struct{}{}
				return

			case <-c.alarmer.Alarm():
				checks = (c.flusher).Flush(ctx, checks)
			}
		}
	}(ctx)
}

// Метод для сохранения сущности
func (c *checkSaverPeriodic) Save(ctx context.Context, check models.Check) error {
	c.checks <- check
	return nil
}

// Закрытие Saver-a: ждем завершения
func (c *checkSaverPeriodic) Close() {
	<-c.done
}

// NewSaver возвращает Saver с поддержкой периодического сохранения
func NewSaver(capacity uint, period time.Duration, alarmer alarmer.Alarmer, flusher flusher.CheckFlusher) Saver {
	// Capacity задается параметром Saver-a
	checks := make(chan models.Check, capacity)
	done := make(chan struct{})

	return &checkSaverPeriodic{
		alarmer: alarmer,
		flusher: flusher,
		checks:  checks,
		done:    done,
	}
}
