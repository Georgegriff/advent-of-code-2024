package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt", 25)
	testutils.ExpectToMatchInt(t, solution, 55312)
}

func TestSolveInput(t *testing.T) {
	solution := Solve("./input.txt", 75)
	testutils.ExpectToMatchInt(t, solution, 221280540398419)
}
