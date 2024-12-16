package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	testutils.ExpectToMatchInt(t, solution, 9021)
}

func TestSolveEdgeCase(t *testing.T) {
	solution := Solve("./edge-case-1.txt")
	testutils.ExpectToMatchInt(t, solution, 406)
}

func TestSolveEdgeCase2(t *testing.T) {
	solution := Solve("./edge-case-2.txt")
	testutils.ExpectToMatchInt(t, solution, 509)
}

func TestSolveEdgeCase4(t *testing.T) {
	solution := Solve("./edge-case-4.txt")
	testutils.ExpectToMatchInt(t, solution, 1216)
}

func TestSolveEdgeCase5(t *testing.T) {
	solution := Solve("./edge-case-5.txt")
	testutils.ExpectToMatchInt(t, solution, 1020)
}
