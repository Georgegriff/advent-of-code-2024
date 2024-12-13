package main

import (
	"aoc/src/aoc/day_13/equations"
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"regexp"
)

func Solve(
	path string,
	adjustment int,
) int {
	file := readfile.Open(path)
	defer file.Close()

	a1Coefficient := 0
	b1Coefficient := 0

	a2Coefficient := 0
	b2Coefficient := 0
	counter := 0

	sum := 0

	err := readfile.ReadLine(file, func(line string) error {
		if line == "" {
			return nil
		}
		if counter == 0 || counter == 1 {
			re := regexp.MustCompile(`Button [A|B]: X\+([0-9]+), Y\+([0-9]+)`)
			matches := re.FindAllStringSubmatch(line, -1)
			if len(matches) != 0 {
				match := matches[0]
				if counter == 0 {
					a1Coefficient = utils.ToInt(match[1])
					a2Coefficient = utils.ToInt(match[2])
				} else if counter == 1 {
					b1Coefficient = utils.ToInt(match[1])
					b2Coefficient = utils.ToInt(match[2])
				}
			}
			counter++
		} else if counter == 2 {
			counter = 0
			re2 := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
			matches := re2.FindAllStringSubmatch(line, -1)
			if len(matches) != 0 {
				match := matches[0]
				a1Output := utils.ToInt(match[1])
				a2Output := utils.ToInt(match[2])
				eq1 := equations.MakeEquation(a1Output, a1Coefficient, b1Coefficient, adjustment)
				eq2 := equations.MakeEquation(a2Output, a2Coefficient, b2Coefficient, adjustment)
				if eq1.SolveEquations(eq2) {
					sum += eq2.GetCost(3, 1)
				}

			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return sum
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt", 10000000000000))
}
