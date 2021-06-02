package saver

import (
	"sync"
	"time"

	"github.com/ozoncp/ocp-check-api/internal/flusher"
	"github.com/ozoncp/ocp-check-api/internal/models"
)

type Saver interface {
	Save(entity models.Check)
	Init()
	Close()
}

// Saver с поддержкой периодического сохранения
type checkSaverPeriodic struct {
	capacity uint
	entities []models.Check
	period   time.Duration
	flusher  *flusher.CheckFlusher
	ticker   *time.Ticker
	mtx      *sync.Mutex
	closeCh  chan bool
	wg       *sync.WaitGroup
}

func (c *checkSaverPeriodic) Init() {
	// Capacity задается параметром Saver-a
	c.entities = make([]models.Check, 0, c.capacity)

	c.mtx = &sync.Mutex{}

	c.wg = &sync.WaitGroup{}
	// Группа wg имеет одно задание: Flush на завершении
	c.wg.Add(1)

	c.ticker = time.NewTicker(c.period)
	c.closeCh = make(chan bool)

	// Flush на закрытие или по тику
	go func() {
		for {
			select {
			case <-c.closeCh:
				c.mtx.Lock()
				c.entities = (*c.flusher).Flush(c.entities)
				c.mtx.Unlock()
				c.wg.Done()
				return
			case <-c.ticker.C:
				c.mtx.Lock()
				c.entities = (*c.flusher).Flush(c.entities)
				c.mtx.Unlock()
			}
		}
	}()
}

// Метод для добавления сущности в slice, который защищен мьютексом
func (c *checkSaverPeriodic) Save(entity models.Check) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.entities = append(c.entities, entity)
}

// Закрытие Saver-a: останавливаем тикер, передаем в канал о закрытии и ждем завершения
func (c *checkSaverPeriodic) Close() {
	c.ticker.Stop()
	c.closeCh <- true
	c.wg.Wait()
}

// NewSaver возвращает Saver с поддержкой периодического сохранения
func NewSaver(capacity uint, period time.Duration, flusher flusher.CheckFlusher) Saver {
	return &checkSaverPeriodic{capacity, make([]models.Check, capacity), period, &flusher, nil, nil, nil, nil}
}
