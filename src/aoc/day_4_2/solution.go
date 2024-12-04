package main

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
)

func main() {
	// file := readfile.Open("./test.txt")
	file := readfile.Open("./input.txt")
	defer file.Close()
	sum := 0
	var grid [][]string
	err := readfile.ReadLine(file, func(line string) error {
		row := make([]string, len(line))
		for i, char := range line {
			row[i] = string(char)
		}
		grid = append(grid, row)
		return nil
	})

	for y, row := range grid {
		for x := range row {
			coordinate := GridCoordinate{x, y}
			sum += coordinate.GetMASTotal(grid)
		}
	}
	fmt.Printf("The answer is %v\n", sum)

	if err != nil {
		log.Fatal(err)
	}
}
