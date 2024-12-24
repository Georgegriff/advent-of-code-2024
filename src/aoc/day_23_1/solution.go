package main

import (
	"aoc/src/aoc/day_23_1/computer"
	"fmt"
)

func Solve(
	path string,

) int {
	network := computer.LoadNetworkConnections(path)
	trios := computer.FindNetworks("t", network)

	return len(trios)
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
