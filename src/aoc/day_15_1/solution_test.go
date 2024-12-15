package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolveSimple(t *testing.T) {
	solution := Solve("./simple.txt")
	testutils.ExpectToMatchInt(t, solution, 2028)
}
func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	testutils.ExpectToMatchInt(t, solution, 10092)
}
