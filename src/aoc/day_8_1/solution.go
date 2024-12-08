package main

import (
	"aoc/src/aoc/day_8_1/grid"
	"fmt"
)

func Solve(
	path string,
) int {
	grid := grid.LoadGrid(path)
	fmt.Println(grid)

	for antennaType, coordinatesForType := range grid.Antennas {
		fmt.Printf("\nChecking antennas: %s", antennaType)
		for i := 0; i < len(coordinatesForType); i++ {
			for j := 0; j < len(coordinatesForType); j++ {
				// don't check self
				if i != j {
					coord := coordinatesForType[i]
					other := coordinatesForType[j]
					points := coord.CalculateAntinodePoints(other)
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
