package models

import (
	"fmt"
)

type Check struct {
	ID         uint64 `json:"id"`
	SolutionID uint64 `json:"solutionID,omitempty"`
	TestID     uint64 `json:"testID,omitempty"`
	RunnerID   uint64 `json:"runnerID,omitempty"`
	Success    bool   `json:"success,omitempty"`
}

func NewCheck(checkID uint64, solutionID uint64, testID uint64, runnerID uint64, success bool) (*Check, error) {
	check := &Check{
		ID:         checkID,
		SolutionID: solutionID,
		TestID:     testID,
		RunnerID:   runnerID,
		Success:    success}
	return check, nil
}

func (c *Check) String() string { return fmt.Sprintf("Check(id=%v)", c.ID) }
