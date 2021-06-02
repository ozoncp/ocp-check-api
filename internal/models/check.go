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

func (c *Check) String() string { return fmt.Sprintf("Check(id=%v)", c.ID) }
