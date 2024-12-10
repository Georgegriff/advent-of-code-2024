package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"math"
	"strings"
)

func main() {
	file := readfile.Open("./input.txt")
	defer file.Close()
	safeCounter := 0
	err := readfile.ReadLine(file, func(line string) error {
		var numPrev int
		isSafe := true

		var direction string
		numStrs := strings.Split(line, " ")
		for i, numStr := range numStrs {

			num := utils.ToInt(numStr)

			if i == 0 {
				numPrev = num
				continue
			} else {
				if numPrev < num {
					if direction == "" {
						direction = "increasing"
					} else if direction == "decreasing" {
						isSafe = false
						break
					}
				} else if numPrev > num {
					if direction == "" {
						direction = "decreasing"
					} else if direction == "increasing" {
						isSafe = false
						break
					}
				} else {
					isSafe = false
					break
				}

			}
			diff := math.Abs(float64(numPrev) - float64(num))
			if diff > 3 {
				isSafe = false
				break
			}
			numPrev = num
		}
		if !isSafe {
			fmt.Printf("line: %v is UNSAFE\n", line)
		} else {

			fmt.Printf("line: %v is SAFE\n", line)
			safeCounter++
		}

		return nil
	})

	fmt.Printf("The answer is %v\n", safeCounter)

	if err != nil {
		log.Fatal(err)
	}
}
