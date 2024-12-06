package main

import (
	"fmt"
)

func Solve(
	path string,
) int {

	puzzleMap := LoadMap(path)
	guard := puzzleMap.Guard
	guard.CompletePatrol(puzzleMap)

	return len(guard.VisitedCoordinates)

}

func main() {
	fmt.Printf("The answer is %v\n", Solve("./input.txt"))
}
