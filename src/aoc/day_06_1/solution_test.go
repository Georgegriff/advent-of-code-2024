package main

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"regexp"
	"testing"
)

func checkState(t *testing.T, testInput string, currentState *Map) {
	printedMap := fmt.Sprint(currentState)
	if testInput != printedMap {
		t.Errorf(`
Expected:
%s
Actual:
%s
`, testInput, printedMap)
	}

}

func expectToMatch(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("got %d, wanted %d", actual, expected)
	}
}

func expectedLocationsCount(input string) int {
	re := regexp.MustCompile("X")
	matches := re.FindAllStringIndex(input, -1)
	return len(matches)
}
func TestMovement(t *testing.T) {
	testInput := readfile.ReadFileToString("./test.txt")

	puzzleMap := LoadMap("./test.txt")
	// check initial
	checkState(t, testInput, puzzleMap)

	// move 1
	guard := puzzleMap.Guard
	guard.MoveToNextObstacle(puzzleMap)
	move := readfile.ReadFileToString("./states/move-1.txt")
	checkState(t, move, puzzleMap)

	expectToMatch(t, len(guard.VisitedCoordinates), expectedLocationsCount(move))

	// move 2
	move = readfile.ReadFileToString("./states/move-2.txt")
	guard.TurnRight()
	guard.MoveToNextObstacle(puzzleMap)
	checkState(t, move, puzzleMap)

	expectToMatch(t, len(guard.VisitedCoordinates), expectedLocationsCount(move))
}

func TestPatrolMapCompleted(t *testing.T) {
	testInput := readfile.ReadFileToString("./test.txt")
	puzzleMap := LoadMap("./test.txt")

	// check initial
	checkState(t, testInput, puzzleMap)

	completedPatrol := readfile.ReadFileToString("./states/completed-patrol.txt")

	guard := puzzleMap.Guard

	guard.CompletePatrol(puzzleMap)
	checkState(t, completedPatrol, puzzleMap)

	expectToMatch(t, len(guard.VisitedCoordinates), expectedLocationsCount(completedPatrol))

}

func TestSolve(t *testing.T) {
	solution := Solve("./test.txt")
	expected := 41

	if solution != expected {
		t.Errorf("got %d, wanted %d", solution, expected)

	}
}
