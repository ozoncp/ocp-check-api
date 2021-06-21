// Package producer defines Producer interface and implements kafkaProducer type which
// send Create/Update/Delete events for models.Check and models.Type into Kafka instance
// and corresponding topics "ocp-check" and "ocp-test"
//
package producer

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

const (
	checkTopic = "ocp-check"
	testTopic  = "ocp-test"
	capacity   = 1024
)

type Producer interface {
	SendCheckEvent(event CheckEvent) error
	SendTestEvent(event TestEvent) error
}

// NewProducer creates instance of kafkaProducer
func NewProducer(ctx context.Context, brokers []string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	messages := make(chan *sarama.ProducerMessage, capacity)

	p := &kafkaProducer{producer: producer, messages: messages}
	go handleMessages(ctx, p)
	return p, err
}

// String represents type of event as string
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

// String represents type of event as string
func (t *TestEvent) String() string {
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

// prepareCheckMessage serializes CheckEvent into sarama.ProducerMessage pointer.
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

// prepareTestMessage serializes TestEvent into sarama.ProducerMessage pointer.
func prepareTestMessage(topic string, event TestEvent, timestamp time.Time) *sarama.ProducerMessage {
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

type kafkaProducer struct {
	producer sarama.SyncProducer
	messages chan *sarama.ProducerMessage
}

func handleMessages(ctx context.Context, p *kafkaProducer) {
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

// SendCheckEvent passes CheckEvent into messages channel for sending to Kafka
func (p *kafkaProducer) SendCheckEvent(event CheckEvent) error {
	p.messages <- prepareCheckMessage(checkTopic, event, time.Now())
	return nil
}

// SendTestEvent passes TestEvent into messages channel for sending to Kafka
func (p *kafkaProducer) SendTestEvent(event TestEvent) error {
	p.messages <- prepareTestMessage(testTopic, event, time.Now())
	return nil
}
