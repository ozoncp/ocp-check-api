package flusher

import (
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/repo"
	"github.com/ozoncp/ocp-check-api/internal/utils"
)

type CheckFlusher interface {
	Flush(checks []models.Check) []models.Check
}

type checkFlusher struct {
	chunkSize int
	checkRepo repo.CheckRepo
}

func (f *checkFlusher) Flush(checks []models.Check) []models.Check {
	bulks := utils.SplitChecksToBulks(checks, uint(f.chunkSize))

	for i := 0; i < len(bulks); i = i + 1 {
		if err := f.checkRepo.AddChecks(bulks[i]); err != nil {
			return bulks[i]
		}
	}

	return nil
}

func NewCheckFlusher(chunkSize int, checkRepo repo.CheckRepo) CheckFlusher {
	return &checkFlusher{chunkSize: chunkSize, checkRepo: checkRepo}
}

type TestFlusher interface {
	Flush(tests []models.Test) []models.Test
}

type testFlusher struct {
	chunkSize int
	testRepo  repo.TestRepo
}

func (f *testFlusher) Flush(tests []models.Test) []models.Test {
	bulks := utils.SplitTestsToBulks(tests, uint(f.chunkSize))

	for i := 0; i < len(bulks); i = i + 1 {
		if err := f.testRepo.AddTests(bulks[i]); err != nil {
			return bulks[i]
		}
	}
	return nil
}

func NewTestFlusher(chunkSize int, testRepo repo.TestRepo) TestFlusher {
	return &testFlusher{chunkSize: chunkSize, testRepo: testRepo}
}
