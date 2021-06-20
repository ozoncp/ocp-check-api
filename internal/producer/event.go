package producer

import "github.com/ozoncp/ocp-check-api/internal/models"

type EventType = int

const (
	Created EventType = iota
	Updated
	Deleted
)

type CheckEvent struct {
	Type  EventType
	Event models.Check
}

type TestEvent struct {
	Type  EventType
	Event models.Test
}
