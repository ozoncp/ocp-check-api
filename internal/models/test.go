package models

import (
	"fmt"
)

type Test struct {
	ID     uint64 `json:"id"`
	TaskID uint64 `json:"taskID,omitempty"`
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}

func (t *Test) String() string { return fmt.Sprintf("Test(id=%v)", t.ID) }
