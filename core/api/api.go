package api

import "fmt"

type Check struct {
	CheckId    uint64
	SolutionId uint
	TestId     uint
	RunnerId   uint
	Success    bool
}

func (c *Check) Init(id uint64, solutionId uint, testId uint, runnerId uint, success bool) {
	c.CheckId = id
	c.SolutionId = solutionId
	c.TestId = testId
	c.RunnerId = runnerId
	c.Success = success
}

func (c *Check) String() string { return fmt.Sprintf("Check(id=%v)", c.CheckId) }

func (c *Check) Create() {
	fmt.Println("Check: create")
}

func (c *Check) Update() {
	fmt.Println("Check: update")
}

func (c *Check) Delete() {
	fmt.Println("Check: delete")
}
