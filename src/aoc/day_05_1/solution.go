package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"strings"
)

func parseUpdate(
	rules map[string]bool,
	line string,
) Update {
	numbersStr := strings.Split(line, ",")
	values := make([]UpdateValue, len(numbersStr))
	for i, numbStr := range numbersStr {
		values[i] = UpdateValue{Value: utils.ToInt(numbStr)}
	}

	return Update{
		Rules:  rules,
		Values: values,
	}
}

func Solve(
	path string,
) int {
	file := readfile.Open(path)
	defer file.Close()
	sum := 0
	rules := make(map[string]bool)

	processingRules := true
	err := readfile.ReadLine(file, func(line string) error {
		if line == "" {
			processingRules = false
		} else {
			if processingRules {
				rules[line] = true
			} else {
				update := parseUpdate(rules, line)
				if !update.CheckViolations() {
					sum += update.GetMiddle().Value

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
	fmt.Printf("The answer is %v\n", Solve("./input.txt"))
}
