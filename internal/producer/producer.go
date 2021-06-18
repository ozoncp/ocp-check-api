package producer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

var brokers = []string{"127.0.0.1:9092"}

const (
	topic    = "ocp-check"
	capacity = 1024
)

type Producer interface {
	SendEvent(event CheckEvent) error
}

func NewProducer(ctx context.Context) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	messages := make(chan *sarama.ProducerMessage, capacity)

	p := &checkProducer{producer: producer, messages: messages}
	go handleMessages(ctx, p)
	return p, err
}

func (t *CheckEvent) String() string {
	switch t.Type {
	case Created:
		return "created"
	case Updated:
		return "updated"
	case Deleted:
		return "removed"
	default:
		return "unknown"
	}
}

func prepareCheckMessage(topic string, event CheckEvent, timestamp time.Time) *sarama.ProducerMessage {
	b, _ := json.Marshal(event.Event)
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:       sarama.StringEncoder(event.String()),
		Value:     sarama.ByteEncoder(b),
		Headers:   []sarama.RecordHeader{},
		Metadata:  nil,
		Offset:    0,
		Partition: -1,
		Timestamp: timestamp,
	}

	return msg
}

type checkProducer struct {
	producer sarama.SyncProducer
	messages chan *sarama.ProducerMessage
}

func handleMessages(ctx context.Context, p *checkProducer) {
	for {
		select {
		case msg := <-p.messages:
			_, _, err := p.producer.SendMessage(msg)
			if err != nil {
				log.Error().Err(err).Msg("handleMessages error")
			}
		case <-ctx.Done():
			close(p.messages)
			return
		}
	}
}

func (p *checkProducer) SendEvent(event CheckEvent) error {
	p.messages <- prepareCheckMessage(topic, event, time.Now())
	return nil
}
