package main

import "testing"

// go test
// go test -v
func TestSum(t *testing.T) {
	result := Sum(1, 1)
	if result != 2 {
		t.Errorf("1+1 expected %d but received %d", 2, result)
	}
}

func TestMultiple(t *testing.T) {
	result := Multiple(1, 2)
	if result != 2 {
		t.Errorf("1x2 expected %d but received %d", 2, result)
	}
}
