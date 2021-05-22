package utils

import (
	"testing"
)

func TestBatchSlice(t *testing.T) {
	slice := []string{"one", "two", "three"}
	var batchSlice [][]string
	batchSlice, _ = BatchSlice(slice, 2)
	got := len(batchSlice)
	want := 2
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestTransposeMap(t *testing.T) {
	sourceMap := map[string]Any{"key1": "value1", "key2": "value2", "key3": "value2"}
	transposedMap, _ := TransposeMap(sourceMap)
	got := len(transposedMap)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
