package flusher

import (
	"github.com/ozoncp/ocp-check-api/internal/models"
)

type CheckFlusher interface {
	Flush(tasks []models.Check) []models.Check
}

type TestFlusher interface {
	Flush(tasks []models.Test) []models.Test
}
