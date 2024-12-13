package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolveSimple(t *testing.T) {
	solution := Solve("./simple.txt", 0)
	testutils.ExpectToMatchInt(t, solution, 280) // Ensure this is the correct expected value
}

func TestSolveSimpleNoMatch(t *testing.T) {
	solution := Solve("./simple-2.txt", 0)
	testutils.ExpectToMatchInt(t, solution, 0) // Ensure this is the correct expected value
}

func TestSolveTest(t *testing.T) {
	solution := Solve("./test.txt", 0)
	testutils.ExpectToMatchInt(t, solution, 480) // Ensure this is the correct expected value
}

func TestSolve(t *testing.T) {
	solution := Solve("./input.txt", 0)
	testutils.ExpectToMatchInt(t, solution, 31897) // Ensure this is the correct expected value
}

func TestSolvePart2(t *testing.T) {
	solution := Solve("./input.txt", 10000000000000)
	testutils.ExpectToMatchInt(t, solution, 87596249540359) // Ensure this is the correct expected value
}
