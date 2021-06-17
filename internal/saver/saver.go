package saver

import (
	"context"
	"time"

	"github.com/ozoncp/ocp-check-api/internal/alarmer"
	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

type Saver interface {
	Save(entity models.Check) error
	Init() error
	Close()
}

// Saver с поддержкой периодического сохранения
type checkSaverPeriodic struct {
	ctx      context.Context
	alarmer  alarmer.Alarmer
	flusher  flusher.CheckFlusher
	checksCh chan models.Check
	closeCh  chan struct{}
	doneCh   chan struct{}
}

func (c *checkSaverPeriodic) Init() error {
	var entities []models.Check

	if err := c.alarmer.Init(); err != nil {
		return err
	}
	// Flush на закрытие или по тику
	go func() {
		for {
			select {
			case e := <-c.checksCh:
				entities = append(entities, e)
			case <-c.closeCh:
				_ = (c.flusher).Flush(c.ctx, entities)
				c.doneCh <- struct{}{}
				return
			case <-c.alarmer.Alarm():
				entities = (c.flusher).Flush(c.ctx, entities)
			}
		}
	}()

	return nil
}

// Метод для сохранения сущности
func (c *checkSaverPeriodic) Save(check models.Check) error {
	c.checksCh <- check
	return nil
}

// Закрытие Saver-a: останавливаем тикер, передаем в канал о закрытии и ждем завершения
func (c *checkSaverPeriodic) Close() {
	c.alarmer.Close()
	c.closeCh <- struct{}{}
	<-c.doneCh
}

// NewSaver возвращает Saver с поддержкой периодического сохранения
func NewSaver(ctx context.Context, capacity uint, period time.Duration, alarmer alarmer.Alarmer, flusher flusher.CheckFlusher) Saver {
	// Capacity задается параметром Saver-a
	checksCh := make(chan models.Check, capacity)
	closeCh := make(chan struct{})
	doneCh := make(chan struct{})

	return &checkSaverPeriodic{
		ctx:      ctx,
		alarmer:  alarmer,
		flusher:  flusher,
		checksCh: checksCh,
		closeCh:  closeCh,
		doneCh:   doneCh,
	}
}
