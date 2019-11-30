package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test
// go test -v
func TestSum(t *testing.T) {
	result := Sum(1, 1)
	assert.Equal(t, 2, result, "they should be equal")
	// if result != 2 {
	// 	t.Errorf("1+1 expected %d but received %d", 2, result)
	// }
}

func TestMultiple(t *testing.T) {
	result := Multiple(1, 2)
	assert.Equal(t, 2, result, "they should be equal")
}
