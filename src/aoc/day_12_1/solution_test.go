package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolveSimple2(t *testing.T) {
	solution := Solve("./simple-2.txt")
	xSum := 4 * 4
	oSum := 21 * (20 + 16)
	testutils.ExpectToMatchInt(t, solution, xSum+oSum) // Ensure this is the correct expected value
}

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	testutils.ExpectToMatchInt(t, solution, 1930) // Ensure this is the correct expected value
}
