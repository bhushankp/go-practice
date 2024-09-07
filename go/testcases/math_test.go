package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 4)
	expected := 6

	if result != expected {
		t.Errorf("Add(2,4) returned %d expected %d", result, expected)
	} else {
		fmt.Println("test passed!!!")
	}

}
