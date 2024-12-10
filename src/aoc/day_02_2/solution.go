package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"math"
	"strings"
)

func safetyChecker(numStrs []string) bool {
	var numPrev int
	isSafe := true

	var direction string
	hasRemoved := false
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
					if !hasRemoved {
						hasRemoved = true
						// skip over one number that would have made it invalid, don't update numPrev to pretend it doesn't exist
					} else {
						isSafe = false
						break
					}
				}
			} else if numPrev > num {
				if direction == "" {
					direction = "decreasing"
				} else if direction == "increasing" {
					if !hasRemoved {
						hasRemoved = true
						// skip over one number that would have made it invalid, don't update numPrev to pretend it doesn't exist
					} else {
						isSafe = false
						break
					}
				}
			} else {
				if !hasRemoved {
					hasRemoved = true
					// skip over one number that would have made it invalid, don't update numPrev to pretend it doesn't exist
				} else {
					isSafe = false
					break
				}
			}

		}
		diff := math.Abs(float64(numPrev) - float64(num))
		if diff > 3 {
			if !hasRemoved {
				hasRemoved = true
				// skip over one number that would have made it invalid, don't update numPrev to pretend it doesn't exist

			} else {
				isSafe = false
				break
			}
		}
		numPrev = num
	}

	return isSafe
}

func main() {
	file := readfile.Open("./input.txt")
	defer file.Close()
	safeCounter := 0
	err := readfile.ReadLine(file, func(line string) error {

		numStrs := strings.Split(line, " ")
		isSafe := safetyChecker(numStrs)
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
