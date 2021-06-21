// Package saver defines the Saver interface and implements periodic Saver, which flushes
// models.Check after timeout elapsed.
//
package saver

import (
	"context"

	"github.com/ozoncp/ocp-check-api/internal/alarmer"
	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

// Interface which describes API to save models.Check type.
type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, check models.Check) error
	Close()
}

// periodic Saver: supports flush on alarm
type checkSaverPeriodic struct {
	alarmer alarmer.Alarmer
	flusher flusher.CheckFlusher
	checks  chan models.Check
	done    chan struct{}
}

// Init: handler for different kinds of channels
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

// Save: it pushes check into the checks channel.
func (c *checkSaverPeriodic) Save(ctx context.Context, check models.Check) error {
	c.checks <- check
	return nil
}

// Close: wait until all jobs will be completed.
func (c *checkSaverPeriodic) Close() {
	<-c.done
}

// NewSaver: creates and returns periodic Saver, where
// - capacity is a size of beffered channel
// - alarmer is a kind of alarm by timeout
// - flusher is a mechanism for flushing collected checks
func NewSaver(capacity uint, alarmer alarmer.Alarmer, flusher flusher.CheckFlusher) Saver {
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
