package repo

import (
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
