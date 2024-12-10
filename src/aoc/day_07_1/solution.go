package main

import (
	"aoc/src/aoc/day_07_1/equation"
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
	"strings"
)

func Solve(
	path string,
) int {
	file := readfile.Open(path)
	defer file.Close()
	sum := 0

	err := readfile.ReadLine(file, func(line string) error {
		parts := strings.Split(line, ": ")
		target := parts[0]
		input := strings.Split(parts[1], " ")
		equation := equation.MakeEquation(target, input)
		if equation.HasValidEquation() {
			sum += equation.Target
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return sum
}

func main() {
	fmt.Printf("The answer is %v\n", Solve("./input.txt"))
}
