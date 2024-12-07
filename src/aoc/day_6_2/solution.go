package main

import (
	"fmt"
)

func Solve(
	path string,
) int {

	puzzleMap := LoadMap(path)
	var cycles map[*Coordinate]bool = make(map[*Coordinate]bool)
	// startPosition := guard.Position
	// startDirection := guard.Direction

	for i, row := range puzzleMap.coordinates {
		for j := range row {
			puzzleMap = LoadMap(path)
			coord := puzzleMap.coordinates[i][j]
			guard := puzzleMap.Guard
			if !coord.Obstacle && coord.Guard == nil {
				coord.Obstacle = true
				// todo need to prevent cycles now
				completed, err := guard.CompletePatrol(puzzleMap)
				if !completed && err != nil {
					cycles[coord] = true
					puzzleMap.PrintMapState()

				}
				coord.Obstacle = false
			}
			guard.Leave()

		}
	}

	return len(cycles)

}

func main() {
	fmt.Printf("\nThe answer is %d\n", Solve("./test.txt"))
}
