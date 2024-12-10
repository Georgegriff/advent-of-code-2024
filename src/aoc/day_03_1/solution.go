package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"regexp"
)

func main() {
	file := readfile.Open("./input.txt")
	defer file.Close()
	sum := 0
	err := readfile.ReadLine(file, func(line string) error {

		re := regexp.MustCompile(`mul\((?P<Num1>[0-9]{1,3}),(?<Num2>[0-9]{1,3})\)`)
		mults := re.FindAllString(line, -1)
		for _, mult := range mults {
			match := re.FindStringSubmatch(mult)
			fmt.Println(match)
			groupNames := re.SubexpNames()
			fmt.Println(groupNames)
			currSum := 0
			for i := range groupNames {
				if i == 0 {
					// first group is empty
					continue
				}
				// +1 because match contains full match, followed by groups
				numStr := match[i]
				num := utils.ToInt(numStr)
				if i == 1 {
					currSum = num
				} else {
					currSum = currSum * num
				}
			}
			sum += currSum
		}

		return nil
	})

	fmt.Printf("The answer is %v\n", sum)

	if err != nil {
		log.Fatal(err)
	}
}
