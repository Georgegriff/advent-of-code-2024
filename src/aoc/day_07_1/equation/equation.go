package equation

import (
	"aoc/src/aoc/day_07_1/permutation"
	"aoc/src/aoc/utils"
)

var operators = []string{"*", "+"}

type Equation struct {
	Target       int
	permutations []permutation.Permutation
}

func MakeEquation(target string, input []string) *Equation {
	targetInt := utils.ToInt(target)
	permutations := permutation.GeneratePermutations(input, operators)

	return &Equation{
		Target:       targetInt,
		permutations: permutations,
	}
}

func (e *Equation) HasValidEquation() bool {
	for _, permutation := range e.permutations {
		if permutation.SumWithoutPrecedence() == e.Target {
			return true
		}
	}
	return false
}
