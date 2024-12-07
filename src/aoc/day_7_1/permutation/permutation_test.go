package permutation

import (
	"aoc/src/testutils"
	"fmt"
	"math"
	"testing"
)

func calcExpectedPermSize(inputLength int, operatorCount int) int {
	return int(math.Pow(float64(operatorCount), float64(inputLength-1)))
}

func TestGeneratePermutations1(t *testing.T) {
	input := []string{"9"}
	operators := []string{"+", "*"}
	permutations := GeneratePermutations(input, operators)
	expected := []string{
		"9",
	}
	testutils.ExpectToMatchInt(t, len(permutations), calcExpectedPermSize(len(input), len(operators)))
	for i, perm := range permutations {
		testutils.ExpectToMatchString(t, expected[i], fmt.Sprint(perm))
	}
}

func TestGeneratePermutations2(t *testing.T) {
	input := []string{"9", "3"}
	operators := []string{"+", "*"}
	permutations := GeneratePermutations(input, operators)
	expected := []string{
		"9+3", "9*3",
	}
	testutils.ExpectToMatchInt(t, len(permutations), calcExpectedPermSize(len(input), len(operators)))
	for i, perm := range permutations {
		testutils.ExpectToMatchString(t, expected[i], fmt.Sprint(perm))
	}
}
func TestGeneratePermutations3(t *testing.T) {
	input := []string{"9", "5", "7"}
	operators := []string{"+", "*"}
	permutations := GeneratePermutations(input, operators)
	expected := []string{
		"9+5+7", "9*5+7", "9+5*7", "9*5*7",
	}
	testutils.ExpectToMatchInt(t, len(permutations), calcExpectedPermSize(len(input), len(operators)))
	fmt.Println(permutations)
	for i, perm := range permutations {
		testutils.ExpectToMatchString(t, expected[i], fmt.Sprint(perm))
	}
}

func TestGeneratePermutations4(t *testing.T) {
	input := []string{"10", "9", "5", "7"}
	operators := []string{"+", "*"}
	permutations := GeneratePermutations(input, operators)
	fmt.Print(permutations)
	expected := []string{
		"10+9+5+7", "10*9+5+7", "10+9*5+7", "10*9*5+7", "10+9+5*7", "10*9+5*7", "10+9*5*7", "10*9*5*7",
	}

	testutils.ExpectToMatchInt(t, len(permutations), calcExpectedPermSize(len(input), len(operators)))
	for i, perm := range permutations {
		testutils.ExpectToMatchString(t, expected[i], fmt.Sprint(perm))
	}
}

func TestSumWithoutPrecedence(t *testing.T) {
	perm := Permutation{
		representation: []string{"10", "+", "19"},
		operators:      []string{"+", "*"},
	}
	testutils.ExpectToMatchInt(t,
		perm.SumWithoutPrecedence(), 29)

	perm = Permutation{
		representation: []string{"10", "*", "19"},
		operators:      []string{"+", "*"},
	}
	testutils.ExpectToMatchInt(t,
		perm.SumWithoutPrecedence(), 190)

	perm = Permutation{
		representation: []string{"81", "+", "40", "*", "27"},
		operators:      []string{"+", "*"},
	}
	testutils.ExpectToMatchInt(t,
		perm.SumWithoutPrecedence(), 3267)

}
