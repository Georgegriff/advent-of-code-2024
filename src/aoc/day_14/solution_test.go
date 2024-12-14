package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt", 11, 7, 100)
	testutils.ExpectToMatchInt(t, solution, 12)
}
