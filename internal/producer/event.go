package producer

import "github.com/ozoncp/ocp-check-api/internal/models"

type EventType = int

const (
	// Types of event: creation/updation/deletion
	Created EventType = iota
	Updated
	Deleted
)

// CheckEvent is a Kafka event for models.Check
type CheckEvent struct {
	Type  EventType
	Event models.Check
}

// TestEvent is a Kafka event for models.Test
type TestEvent struct {
	Type  EventType
	Event models.Test
}
