package utils

import (
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-check-api/core/api"
)

// Function converts slice to slice of slices, size batchSize.
// Each call to BatchSlice returns also an error, when batchSize is equal to 0.
func BatchSlice(slice []string, batchSize uint) (batches [][]string, err error) {
	if batchSize == 0 {
		batches = nil
		err = errors.New("invalid batch size")
		return
	}

	for batchSize < uint(len(slice)) {
		slice, batches = slice[batchSize:], append(batches, slice[0:batchSize:batchSize])
	}

	batches = append(batches, slice)
	err = nil
	return
}

// Function transposes map having keys of type string and values of type interface{}.
// Transposed map has keys and values of type string.
func TransposeMap(source interface{}) (dest map[string]string, err error) {
	unboxed, ok := source.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("source should be a map[string]interface{}")
	}
	if unboxed == nil {
		return nil, fmt.Errorf("source is nil")
	}

	m := make(map[string]string, len(unboxed))
	for k, v := range unboxed {
		s := fmt.Sprintf("%v", v)
		if _, ok := m[s]; ok {
			panic(fmt.Sprintf("duplicate key/value: %v/%v", k, s))
		}
		m[s] = k
	}

	dest = m
	err = nil
	return
}

// Helper function
func filter(vs []string, f func(string) bool) []string {
	filtered := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Function filters source slice by criteria of absence of the element in exclusion list.
func Filter(source []string, exclusion []string) []string {
	e := make(map[string]struct{}, len(exclusion))
	var exists = struct{}{}
	for _, v := range exclusion {
		e[v] = exists
	}

	f := func(s string) bool {
		_, ok := e[s]
		return !ok
	}

	return filter(source, f)
}

func SplitToBulks(checks []api.Check, batchSize uint) (batches [][]api.Check) {
	for int(batchSize) < len(checks) {
		checks, batches = checks[batchSize:], append(batches, checks[0:batchSize:batchSize])
	}

	batches = append(batches, checks)
	return
}

func ConvertSliceToMap(checks []api.Check) (map[uint64]api.Check, error) {
	m := make(map[uint64]api.Check, len(checks))
	for _, v := range checks {
		m[v.ID] = v
	}

	return m, nil
}
