package utils

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ozoncp/ocp-check-api/core/api"
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
	s := map[string]api.Check{}
	_, err := TransposeMap(s)
	expectedErrorMsg := "source should be a map[string]interface{}"
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, actual: %v", expectedErrorMsg, err)
}

func TestTransposeMap(t *testing.T) {
	sourceMap := map[string]interface{}{"key1": "value1", "key2": "value2", "key3": "value3"}
	transposedMap, _ := TransposeMap(sourceMap)

	actual := len(transposedMap)
	expected := 3
	assert.Equal(t, expected, actual, fmt.Sprintf("Transposed map len should be equal of %v", expected))

	expectedMap := map[string]string{"value1": "key1", "value2": "key2", "value3": "key3"}
	if diff := cmp.Diff(expectedMap, transposedMap); diff != "" {
		t.Errorf("TransposeMap() mismatch (-expected +actual):\n%s", diff)
	}
}

func TestTransposeMapInterfaceValues(t *testing.T) {
	var sourceMap = map[string]interface{}{"one": 2, "three": "four"}
	var transposedMap, _ = TransposeMap(sourceMap)
	expectedMap := map[string]string{"2": "one", "four": "three"}
	if diff := cmp.Diff(expectedMap, transposedMap); diff != "" {
		t.Errorf("TransposeMap() mismatch (-expected +actual):\n%s", diff)
	}
}
func TestTransposeMapPanic(t *testing.T) {
	f := func() {
		sourceMap := map[string]interface{}{"key1": "value1", "key2": "value2", "key3": "value2"}
		TransposeMap(sourceMap)
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

func TestApiObjectAsString(t *testing.T) {
	var check = &api.Check{}
	var id uint64 = 1
	check.Init(id, 2, 3, 4, true)
	assert.Equal(t, check.String(), fmt.Sprintf("Check(id=%v)", id))
}
