package repo

import (
	"fmt"

	"github.com/ozoncp/ocp-check-api/internal/models"
)

type CheckRepo interface {
	AddChecks(checks []models.Check) error
	RemoveCheck(checkId uint64) error
	DescribeCheck(checkId uint64) (*models.Check, error)
	ListChecks(limit, offset uint64) ([]models.Check, error)
}

type TestRepo interface {
	AddTests(tests []models.Test) error
	RemoveTest(testId uint64) error
	DescribeTest(testId uint64) (*models.Test, error)
	ListTests(limit, offset uint64) ([]models.Test, error)
}

type consoleRepo struct {
}

func (r *consoleRepo) ListChecks(limit, offset uint64) ([]models.Check, error) {
	return []models.Check{{ID: 1}}, nil
}

func (r *consoleRepo) DescribeCheck(checkId uint64) (*models.Check, error) {
	return &models.Check{ID: checkId}, nil
}

func (r *consoleRepo) RemoveCheck(checkId uint64) error {
	return nil
}

func (r *consoleRepo) AddChecks(checks []models.Check) error {
	for _, c := range checks {
		fmt.Printf("+ putting %v into repo\n", c.String())
	}

	return nil
}

func NewCheckRepo() CheckRepo {
	return &consoleRepo{}
}
