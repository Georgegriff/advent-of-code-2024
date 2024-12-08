package main

import (
	"aoc/src/aoc/day_8_2/grid"
	"fmt"
)

func Solve(
	path string,
) int {
	grid := grid.LoadGrid(path)
	fmt.Println(grid)

	yMax := len(grid.Coordinates) - 1
	xMax := len(grid.Coordinates[0]) - 1
	for _, coordinatesForType := range grid.Antennas {
		for i := 0; i < len(coordinatesForType); i++ {
			for j := 0; j < len(coordinatesForType); j++ {
				// don't check self
				if i != j {
					coord := coordinatesForType[i]
					other := coordinatesForType[j]
					points := coord.CalculateAntinodePoints(other, yMax, xMax)
					grid.AddAntinodeToCoordinates(points)
				}
			}
		}
	}

	fmt.Println(grid)
	return len(grid.Antinodes)
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
