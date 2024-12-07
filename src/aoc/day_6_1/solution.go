package main

import (
	"fmt"
)

func Solve(
	path string,
) int {

	puzzleMap := LoadMap(path)
	guard := puzzleMap.Guard
	fmt.Println("Guard patrol:")
	fmt.Print(puzzleMap)
	guard.CompletePatrol(puzzleMap)

	return len(guard.VisitedCoordinates)

}

func main() {
	fmt.Printf("\nThe answer is %v\n", Solve("./test.txt"))
}
