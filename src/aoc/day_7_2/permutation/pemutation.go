package permutation

import (
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"strings"
)

const PLUS = "+"
const MULTIPLY = "*"
const CONCATENATION = "||"

type Permutation struct {
	representation []string
	operators      []string
}

func (p Permutation) String() string {
	return strings.Join(p.representation, "")
}

func (p *Permutation) SumWithoutPrecedence() int {
	rep := p.representation
	sum := utils.ToInt(rep[0])
	var previousOperator string
	isTestCmd := fmt.Sprint(p) == "6*8||6*15"
	if isTestCmd {
		fmt.Println("In the test")
	}
	for i := 1; i < len(rep); i++ {
		current := rep[i]
		if current == PLUS || current == MULTIPLY || current == CONCATENATION {
			previousOperator = current
		} else {
			currentNum := utils.ToInt(current)
			if previousOperator == PLUS {
				sum += currentNum
			} else if previousOperator == MULTIPLY {
				sum *= currentNum
			} else if previousOperator == CONCATENATION {
				sumStr := fmt.Sprintf("%d%d", sum, currentNum)
				sum = utils.ToInt(sumStr)
			} else if previousOperator == "" {
				log.Fatalf("invalid permutation %s. No connecting operator found", p)
			} else {
				log.Fatalf("unsupported operator %s", previousOperator)
			}
		}
	}
	return sum
}

func GeneratePermutations(input []string, operators []string) []Permutation {
	if len(input) == 1 {
		// only one item remaining
		perm := []Permutation{{
			representation: []string{input[0]},
			operators:      operators,
		}}
		return perm
	}
	firstValue := input[0]
	remaining := input[1:]

	combinations := GeneratePermutations(remaining, operators)

	results := []Permutation{}
	for _, combination := range combinations {
		for _, operator := range operators {
			newPart := []string{firstValue, operator}
			representation := append(newPart, combination.representation...)
			perm := Permutation{
				representation: representation,
				operators:      operators,
			}
			results = append(results, perm)
		}

	}
	return results
}
