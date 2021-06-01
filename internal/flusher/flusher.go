package flusher

import (
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/ozoncp/ocp-check-api/internal/repo"
)

type CheckFlusher interface {
	Flush(checks []models.Check) []models.Check
}

type checkFlusher struct {
	chunkSize int
	checkRepo repo.CheckRepo
}

func (c *checkFlusher) Flush(checks []models.Check) []models.Check {
	return []models.Check{}
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
	offset := int(0)
	for offset+f.chunkSize < len(tests) {
		if err := f.testRepo.AddTests(tests[offset:f.chunkSize]); err != nil {
			return tests[offset:f.chunkSize]
		}
		offset += f.chunkSize
	}

	if len(tests) == offset {
		return nil
	}

	// add last chunk
	if err := f.testRepo.AddTests(tests[offset:]); err != nil {
		return tests[offset:]
	}

	return nil
}

func NewTestFlusher(chunkSize int, testRepo repo.TestRepo) TestFlusher {
	return &testFlusher{chunkSize: chunkSize, testRepo: testRepo}
}
