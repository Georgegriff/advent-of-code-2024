package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

func main() {
	file := readfile.Open("./input.txt")
	defer file.Close()

	one := []int{}
	two := []int{}

	err := readfile.ReadLine(file, func(line string) error {
		numbers := strings.Split(line, strings.Repeat(" ", 3))
		if len(numbers) != 2 {
			return fmt.Errorf("line did not contain two numbers: %v", line)
		}

		one = append(one, utils.ToInt(numbers[0]))
		two = append(two, utils.ToInt(numbers[1]))

		return nil
	})

	sort.Ints(one)
	sort.Ints(two)

	sum := 0

	for i := 0; i < len(one); i++ {
		diff := int(math.Abs(float64(one[i]) - float64(two[i])))
		sum += diff
	}

	fmt.Printf("The answer is %v", sum)

	if err != nil {
		log.Fatal(err)
	}
}
