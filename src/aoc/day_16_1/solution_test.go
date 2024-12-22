package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	testutils.ExpectToMatchInt(t, solution, 7036)
}

func TestSolve2(t *testing.T) {
	solution := Solve("./test-2.txt")
	testutils.ExpectToMatchInt(t, solution, 11048)
}
