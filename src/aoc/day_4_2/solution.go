package main

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
)

type GridCoordinates struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
}

func getPermutationLetters(
	grid [][]string, topLeft []int, topRight []int, bottomLeft []int, bottomRight []int,
) GridCoordinates {
	return GridCoordinates{
		TopLeft:     grid[topLeft[1]][topLeft[0]],
		TopRight:    grid[topRight[1]][topRight[0]],
		BottomLeft:  grid[bottomLeft[1]][bottomLeft[0]],
		BottomRight: grid[bottomRight[1]][bottomRight[0]],
	}
}

func getOffsetCoordinate(currX int, currY int, xOffset int, yOffset int) []int {
	coord := []int{
		currX + xOffset,
		currY + yOffset,
	}
	return coord
}

func checkCoordinate(grid [][]string, x int, y int) int {
	currentLetter := grid[y][x]

	if currentLetter != "A" {
		return 0
	}

	columnLength := len(grid)
	rowLength := len(grid[y])
	if x-1 < 0 || y-1 < 0 || y+1 > columnLength-1 || x+1 > rowLength-1 {

		return 0
	}

	topLeft := getOffsetCoordinate(x, y, -1, -1)
	topRight := getOffsetCoordinate(x, y, 1, -1)
	bottomLeft := getOffsetCoordinate(x, y, -1, 1)
	bottomRight := getOffsetCoordinate(x, y, 1, 1)
	letters := getPermutationLetters(grid,
		topLeft, topRight, bottomLeft, bottomRight,
	)
	// Permutations

	// S.S
	// .A.
	// M.M
	if letters.TopLeft == "S" && letters.TopRight == "S" && letters.BottomLeft == "M" && letters.BottomRight == "M" {
		return 1
	}

	// M.M
	// .A.
	// S.S
	if letters.TopLeft == "M" && letters.TopRight == "M" && letters.BottomLeft == "S" && letters.BottomRight == "S" {
		return 1
	}

	// S.M
	// .A.
	// S.M
	if letters.TopLeft == "S" && letters.TopRight == "M" && letters.BottomLeft == "S" && letters.BottomRight == "M" {
		return 1
	}

	// M.S
	// .A.
	// M.S
	if letters.TopLeft == "M" && letters.TopRight == "S" && letters.BottomLeft == "M" && letters.BottomRight == "S" {
		return 1
	}
	return 0
}

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
			sum += checkCoordinate(grid, x, y)
		}
	}
	fmt.Printf("The answer is %v\n", sum)

	if err != nil {
		log.Fatal(err)
	}
}
