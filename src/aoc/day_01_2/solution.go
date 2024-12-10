package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"strings"
)

func main() {

	file := readfile.Open("./input.txt")
	defer file.Close()

	leftHistogram := map[int]int{}
	histogram := map[int]int{}

	err := readfile.ReadLine(file, func(line string) error {
		numbers := strings.Split(line, strings.Repeat(" ", 3))
		if len(numbers) != 2 {
			return fmt.Errorf("line did not contain two numbers: %v", line)
		}

		left := utils.ToInt(numbers[0])
		right := utils.ToInt(numbers[1])

		leftValueCount, ok := leftHistogram[left]
		if ok {
			leftHistogram[left] = leftValueCount + 1
		} else {
			leftHistogram[left] = 1
		}

		rightValCount, ok := histogram[right]
		if ok {
			histogram[right] = rightValCount + 1
		} else {
			histogram[right] = 1
		}

		return nil
	})

	sum := 0
	for k, leftCount := range leftHistogram {
		count, ok := histogram[k]
		if ok {
			sum += (count * k) * leftCount
		}
	}

	fmt.Printf("The answer is: %v", sum)

	if err != nil {
		log.Fatal(err)
	}
}
