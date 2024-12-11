package main

import (
	"aoc/src/aoc/day_11/evolution"
	"aoc/src/aoc/readfile"
	"fmt"
)

func Solve(
	path string,
) int {
	input := readfile.ReadFileToString(path)
	iterations := 75
	evolutionState := evolution.Evolution{
		Stones: input,
	}
	for i := 0; i < iterations; i++ {
		evolutionState = evolutionState.Evolve()
	}

	return evolutionState.CountStones()

}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
