package main

import (
	"aoc/src/aoc/day_14/robot"
	"fmt"
)

func main() {
	lobby := robot.LoadLobby("./input.txt", 101, 103)
	for i := 0; i < 50000; i++ {
		lobby.MoveRobots(1)

		if lobby.AllUnique() {
			fmt.Printf("\n Found A Christmas Tree after: %d seconds", i)

			fmt.Printf("\n%#v\n", lobby)
			break
		}
	}
}
