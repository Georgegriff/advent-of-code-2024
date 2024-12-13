package main

import (
	"aoc/src/aoc/day_12_1/garden"
	"fmt"
)

func Solve(
	path string,
) int {

	garden := garden.LoadGarden(path)

	return garden.GetPrice()
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
