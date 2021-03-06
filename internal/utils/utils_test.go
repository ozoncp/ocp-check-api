package utils

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ozoncp/ocp-check-api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestBatchSliceZeroSize(t *testing.T) {
	const expectedErrorMsg = "invalid batch size"
	var err error
	_, err = BatchSlice([]string{}, 0)
	if err == nil {
		t.Errorf("error expected")
	}
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, actual: %v", expectedErrorMsg, err)
}

func TestBatchSlice(t *testing.T) {
	slice := []string{"one", "two", "three"}
	var batchSlice [][]string
	batchSlice, _ = BatchSlice(slice, 2)
	actual := len(batchSlice)
	expected := 2
	assert.Equal(t, expected, actual, fmt.Sprintf("Resulted slice len should be equal of %v", expected))
}

func TestBatchSliceOneBatchOnly(t *testing.T) {
	var slice = []string{"one", "two", "three"}
	const batchSize = 3
	var batch, _ = BatchSlice(slice, batchSize)
	actual := len(batch)
	expected := 1
	assert.Equal(t, expected, actual, fmt.Sprintf("Resulted slice len should be equal of %v", expected))
}
func TestTransposeNilMap(t *testing.T) {
	var m map[string]interface{} = nil
	_, err := TransposeMap(m)
	expectedErrorMsg := "source is nil"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, actual: %v", expectedErrorMsg, err)
}

func TestTransposeNotMap(t *testing.T) {
	s := map[string]models.Check{}
	_, err := TransposeMap(s)
	expectedErrorMsg := "source should be a map[string]interface{}"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, actual: %v", expectedErrorMsg, err)
}

func TestTransposeMap(t *testing.T) {
	sourceMap := map[string]interface{}{"key1": "value1", "key2": "value2", "key3": "value3"}
	transposedMap, err := TransposeMap(sourceMap)
	if assert.NoError(t, err) {
		actual := len(transposedMap)
		expected := 3
		assert.Equal(t, expected, actual, fmt.Sprintf("Transposed map len should be equal of %v", expected))

		expectedMap := map[string]string{"value1": "key1", "value2": "key2", "value3": "key3"}
		if diff := cmp.Diff(expectedMap, transposedMap); diff != "" {
			t.Errorf("TransposeMap() mismatch (-expected +actual):\n%s", diff)
		}
	}
}

func TestTransposeMapInterfaceValues(t *testing.T) {
	var sourceMap = map[string]interface{}{"one": 2, "three": "four"}
	var transposedMap, err = TransposeMap(sourceMap)
	if assert.NoError(t, err) {
		expectedMap := map[string]string{"2": "one", "four": "three"}
		if diff := cmp.Diff(expectedMap, transposedMap); diff != "" {
			t.Errorf("TransposeMap() mismatch (-expected +actual):\n%s", diff)
		}
	}
}
func TestTransposeMapPanic(t *testing.T) {
	f := func() {
		sourceMap := map[string]interface{}{"key1": "value1", "key2": "value2", "key3": "value2"}
		_, _ = TransposeMap(sourceMap)
	}
	expectedErrorMsg := "duplicate key/value: key3/value2"
	assert.PanicsWithValue(t, expectedErrorMsg, f)
}

func TestFilter(t *testing.T) {
	var filtered = Filter([]string{"one", "two", "three"}, []string{"four", "five", "two"})
	expected := []string{"one", "three"}
	if diff := cmp.Diff(expected, filtered); diff != "" {
		t.Errorf("Filter() mismatch (-expected +actual):\n%s", diff)
	}
}

func TestModelsCheckAsString(t *testing.T) {
	var id uint64 = 1
	var check = &models.Check{ID: id, SolutionID: 2, TestID: 3, RunnerID: 4, Success: true}
	assert.Equal(t, check.String(), fmt.Sprintf("Check(id=%v)", id))
}

func TestModelsTestAsString(t *testing.T) {
	var id uint64 = 10000000000
	var test = &models.Test{ID: id, TaskID: 2, Input: "some input", Output: "panic"}
	assert.Equal(t, test.String(), fmt.Sprintf("Test(id=%v)", id))
}
