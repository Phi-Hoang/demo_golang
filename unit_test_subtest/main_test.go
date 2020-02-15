package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://github.com/golang/go/wiki/TableDrivenTests
type TestCase struct {
	Name     string
	Input1   int
	Input2   int
	Expected int
}

// go test
// go test -v
func TestSum(t *testing.T) {
	tests := []TestCase{{"1+1", 1, 1, 3}, {"1+2", 1, 2, 3}, {"2+3", 2, 3, 5}}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result := Sum(tc.Input1, tc.Input2)
			assert.Equal(t, tc.Expected, result, "they should be equal")
		})
	}
}

func TestMultiple(t *testing.T) {
	tests := []TestCase{{"1x1", 1, 1, 1}, {"1x2", 1, 2, 2}, {"2x3", 2, 3, 6}}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result := Multiple(tc.Input1, tc.Input2)
			assert.Equal(t, tc.Expected, result, "they should be equal")
		})
	}
}
