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

func NewTest(testID uint64, taskID uint64, input string, output string) (*Test, error) {
	test := &Test{
		ID:     testID,
		TaskID: taskID,
		Input:  input,
		Output: output}
	return test, nil
}

func (t *Test) String() string { return fmt.Sprintf("Test(id=%v)", t.ID) }
