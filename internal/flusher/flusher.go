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

type checkFlusher struct {
	chunkSize int
	checkRepo repo.CheckRepo
}

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

func NewCheckFlusher(chunkSize int, checkRepo repo.CheckRepo) CheckFlusher {
	return &checkFlusher{chunkSize: chunkSize, checkRepo: checkRepo}
}

type TestFlusher interface {
	Flush(ctx context.Context, tests []models.Test) []models.Test
}

type testFlusher struct {
	chunkSize int
	testRepo  repo.TestRepo
}

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

func NewTestFlusher(chunkSize int, testRepo repo.TestRepo) TestFlusher {
	return &testFlusher{chunkSize: chunkSize, testRepo: testRepo}
}
