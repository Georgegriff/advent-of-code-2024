package main

import (
	"aoc/src/aoc/day_11/evolution"
	"aoc/src/aoc/readfile"
	"fmt"
)

func Solve(
	path string,
	iterations int,
) int {
	input := readfile.ReadFileToString(path)
	initialState := evolution.MakeEvolution(input, nil)

	return initialState.Evolve(iterations).CountStones()

}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt", 75))
}
