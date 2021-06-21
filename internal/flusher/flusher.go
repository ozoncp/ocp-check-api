// Package flusher defines CheckFlusher and TestFlusher interfaces and
// implements checkFlusher and testFlusher types which
// flushes slice of models.Check and models.Type into repo.
//
package flusher

import (
	"context"

	"github.com/ozoncp/ocp-check-api/internal/models"
	repo "github.com/ozoncp/ocp-check-api/internal/repo"
	"github.com/ozoncp/ocp-check-api/internal/utils"
)

type CheckFlusher interface {
	Flush(ctx context.Context, checks []models.Check) []models.Check
}

// CheckFlusher implementation
type checkFlusher struct {
	chunkSize int
	checkRepo repo.CheckRepo
}

// Flush splits slice of models.Check into batches and sends each batch into repo.
// If error occurred, it returns remained (not flushed) checks.
func (f *checkFlusher) Flush(ctx context.Context, checks []models.Check) []models.Check {
	bulks, err := utils.SplitChecksToBulks(checks, uint(f.chunkSize))
	if err != nil {
		return checks
	}

	for i := 0; i < len(bulks); i = i + 1 {
		if _, err := f.checkRepo.MultiCreateCheck(ctx, bulks[i]); err != nil {
			return checks[i*f.chunkSize:]
		}
	}

	return nil
}

// NewCheckFlusher creates instance of checkFlusher
func NewCheckFlusher(chunkSize int, checkRepo repo.CheckRepo) CheckFlusher {
	return &checkFlusher{chunkSize: chunkSize, checkRepo: checkRepo}
}

type TestFlusher interface {
	Flush(ctx context.Context, tests []models.Test) []models.Test
}

// TestFlusher implementation
type testFlusher struct {
	chunkSize int
	testRepo  repo.TestRepo
}

// Flush splits slice of models.Test into batches and sends each batch into repo.
// If error occurred, it returns remained (not flushed) tests.
func (f *testFlusher) Flush(ctx context.Context, tests []models.Test) []models.Test {
	bulks, err := utils.SplitTestsToBulks(tests, uint(f.chunkSize))
	if err != nil {
		return tests
	}

	for i := 0; i < len(bulks); i = i + 1 {
		if _, err := f.testRepo.MultiCreateTest(ctx, bulks[i]); err != nil {
			return tests[i*f.chunkSize:]
		}
	}
	return nil
}

// NewTestFlusher creates instance of testFlusher
func NewTestFlusher(chunkSize int, testRepo repo.TestRepo) TestFlusher {
	return &testFlusher{chunkSize: chunkSize, testRepo: testRepo}
}
