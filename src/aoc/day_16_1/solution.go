package main

import (
	"aoc/src/aoc/day_16_1/maze"
	"fmt"
)

func Solve(
	path string,

) int {
	maze := maze.LoadMaze(path, maze.EAST)

	return maze.GetBestPathScore()
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
