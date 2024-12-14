package main

import (
	"aoc/src/aoc/day_14/robot"
	"fmt"
)

func Solve(
	path string,
	xSize int,
	ySize int,
	time int,
) int {
	lobby := robot.LoadLobby(path, xSize, ySize)
	lobby.MoveRobots(time)
	return lobby.CalculateSafetyFactor()
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt", 101, 103, 100))
}
