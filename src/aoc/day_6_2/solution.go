package main

import (
	"fmt"
)

func Solve(
	path string,
) int {

	puzzleMap := LoadMap(path)
	guard := puzzleMap.Guard
	var cycles map[*Coordinate]bool = make(map[*Coordinate]bool)
	startPosition := guard.Position
	startDirection := guard.Direction

	for _, row := range puzzleMap.coordinates {
		for _, coord := range row {
			guard := NewGuard(startPosition, startDirection)
			if !coord.Obstacle && coord.Guard == nil {
				coord.Obstacle = true
				// todo need to prevent cycles now
				completed, err := guard.CompletePatrol(puzzleMap)
				if !completed && err != nil {
					cycles[coord] = true
				}
				coord.Obstacle = false
			}
			guard.Leave()

		}
	}

	return len(cycles)

}

func main() {
	fmt.Printf("The answer is %v\n", Solve("./input.txt"))
}
