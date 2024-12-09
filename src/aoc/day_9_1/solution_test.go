package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	testutils.ExpectToMatchInt(t, solution, 1928)

}
