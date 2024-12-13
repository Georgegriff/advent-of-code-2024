package main

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolveSimple(t *testing.T) {
	solution := Solve("./simple.txt") //

	testutils.ExpectToMatchInt(t, solution, 80) // Ensure this is the correct expected value
}

func TestSolveSimple2(t *testing.T) {
	solution := Solve("./simple-2.txt") // OXO

	testutils.ExpectToMatchInt(t, solution, 436) // Ensure this is the correct expected value
}

func TestSolveSimple3(t *testing.T) {
	solution := Solve("./simple-3.txt") // E

	testutils.ExpectToMatchInt(t, solution, 236) // Ensure this is the correct expected value
}

func TestSolveSimple4(t *testing.T) {
	solution := Solve("./simple-4.txt")

	testutils.ExpectToMatchInt(t, solution, 368) // Ensure this is the correct expected value
}

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	testutils.ExpectToMatchInt(t, solution, 1206) // Ensure this is the correct expected value
}

func TestInput(t *testing.T) {
	solution := Solve("./input.txt")
	testutils.ExpectToMatchInt(t, solution, 808796) // Ensure this is the correct expected value
}
