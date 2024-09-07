package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 4)
	expected := 6

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
