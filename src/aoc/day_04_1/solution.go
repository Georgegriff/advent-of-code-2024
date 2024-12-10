package main

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
)

const WORD = "XMAS"

func checkWord(word string) bool {
	return word == WORD
}

func buildWord(grid [][]string, coords [][]int) (string, error) {
	word := ""
	for _, coord := range coords {
		column := coord[0]
		row := coord[1]
		if row < 0 || row > len(grid)-1 {
			return "", fmt.Errorf("coordinate %d out of range of rows", coord)
		}
		gridColumn := grid[row]
		if column < 0 || column > len(gridColumn)-1 {
			return "", fmt.Errorf("coordinate %d out of range of column", coord)
		}
		word += gridColumn[column]
	}
	return word, nil
}

func getWordCount(grid [][]string, coords [][]int) int {
	word, err := buildWord(grid, coords)
	if err == nil {
		if checkWord(word) {
			return 1
		}
	}
	return 0
}

func checkCoordinate(grid [][]string, x int, y int) int {
	wordCount := 0

	offset := len(WORD)
	// 1. →
	var coords [][]int
	for i := 0; i < offset; i++ {
		coord := []int{x + i, y}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 2. ←
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x - i, y}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 3. ↑
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x, y + i}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 4. ↓
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x, y - i}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 5. ↘
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x + i, y + i}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 6. ↖
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x - i, y - i}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 7. ↗
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x + i, y - i}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	// 8. ↙
	coords = nil
	for i := 0; i < offset; i++ {
		coord := []int{x - i, y + i}
		coords = append(coords, coord)
	}
	wordCount += getWordCount(grid, coords)

	return wordCount
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
