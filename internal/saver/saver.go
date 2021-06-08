package saver

import (
	"time"

	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

type Saver interface {
	Save(entity models.Check) error
	Init()
	Close()
}

// Saver с поддержкой периодического сохранения
type checkSaverPeriodic struct {
	flusher  flusher.CheckFlusher
	checksCh chan models.Check
	closeCh  chan bool
	ticker   *time.Ticker
}

func (c *checkSaverPeriodic) Init() {
	var entities []models.Check

	// Flush на закрытие или по тику
	go func() {
		for {
			select {
			case e := <-c.checksCh:
				entities = append(entities, e)
			case <-c.closeCh:
				_ = (c.flusher).Flush(entities)
				return
			case <-c.ticker.C:
				entities = (c.flusher).Flush(entities)
			}
		}
	}()
}

// Метод для сохранения сущности
func (c *checkSaverPeriodic) Save(check models.Check) error {
	c.checksCh <- check
	return nil
}

// Закрытие Saver-a: останавливаем тикер, передаем в канал о закрытии и ждем завершения
func (c *checkSaverPeriodic) Close() {
	c.ticker.Stop()
	c.closeCh <- true
}

// NewSaver возвращает Saver с поддержкой периодического сохранения
func NewSaver(capacity uint, period time.Duration, flusher flusher.CheckFlusher) Saver {
	// Capacity задается параметром Saver-a
	checksCh := make(chan models.Check, capacity)
	closeCh := make(chan bool)
	ticker := time.NewTicker(period)

	return &checkSaverPeriodic{flusher, checksCh, closeCh, ticker}
}
