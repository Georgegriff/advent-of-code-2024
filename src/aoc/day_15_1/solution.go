package main

import (
	"aoc/src/aoc/day_15_1/warehouse"
	"fmt"
)

func Solve(
	path string,

) int {
	warehouse := warehouse.LoadWarehouse(path)

	warehouse.Process()
	return warehouse.GetGPSSum()
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
