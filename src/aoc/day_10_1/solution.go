package main

import (
	topmap "aoc/src/aoc/day_10_1/top_map"
	"fmt"
)

func Solve(
	path string,
) int {
	tMap := topmap.LoadMap(path)
	return tMap.GetTrailHeadSum()
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
