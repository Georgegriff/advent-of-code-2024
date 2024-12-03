package main

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"regexp"
)

const NUM_1 = "Num1"
const NUM_2 = "Num2"
const DO = "Do"
const DONT = "Dont"

func getGroups(re *regexp.Regexp, text string) []map[string]string {

	// Find all matches
	matches := re.FindAllStringSubmatch(text, -1)
	groupNames := re.SubexpNames()

	// Collect results
	var results []map[string]string
	for _, match := range matches {
		matchMap := make(map[string]string)
		for i, value := range match {
			if i == 0 || value == "" { // Skip the full match and empty groups
				continue
			}
			matchMap[groupNames[i]] = value
		}
		results = append(results, matchMap)
	}
	return results
}

func main() {
	// file := readfile.Open("./test.txt")
	file := readfile.Open("./input.txt")
	defer file.Close()
	sum := 0
	isAllowed := true
	err := readfile.ReadLine(file, func(line string) error {

		re := regexp.MustCompile(`mul\((?P<Num1>[0-9]{1,3}),(?<Num2>[0-9]{1,3})\)|(?P<Dont>don't\(\))|(?P<Do>do\(\))`)
		groups := getGroups(re, line)
		for _, match := range groups {
			fmt.Println(match)

			// number match
			if len(match) == 2 {
				if isAllowed {
					numOne := utils.ToInt(match[NUM_1])
					numTwo := utils.ToInt(match[NUM_2])
					sum = sum + (numOne * numTwo)
					fmt.Println("Allowed")
				}
				continue
			} else if len(match) == 1 {
				if match[DO] != "" {
					isAllowed = true
					fmt.Println("Resetting allowed")
					continue
				} else if match[DONT] != "" {
					isAllowed = false
					fmt.Println("Disallowing")
					continue
				}
			}
			log.Fatalf("no numbers or instructions found %v", match)

		}

		return nil
	})

	fmt.Printf("The answer is %v\n", sum)

	if err != nil {
		log.Fatal(err)
	}
}
