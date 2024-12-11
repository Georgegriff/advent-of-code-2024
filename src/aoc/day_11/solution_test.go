package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt", 25)
	testutils.ExpectToMatchInt(t, solution, 55312)
}
