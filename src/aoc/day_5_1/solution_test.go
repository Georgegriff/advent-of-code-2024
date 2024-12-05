package main

import "testing"

func TestInput(t *testing.T) {
	solution := Solve("./test.txt")
	expected := 143

	if solution != expected {
		t.Errorf("got %d, wanted %d", solution, expected)

	}
}
