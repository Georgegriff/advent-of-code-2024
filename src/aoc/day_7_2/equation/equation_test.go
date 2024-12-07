package equation

import (
	"aoc/src/testutils"
	"testing"
)

func TestHasValidEquation(t *testing.T) {
	equation := MakeEquation("190", []string{"10", "19"})
	testutils.ExpectToMatchBool(t, equation.HasValidEquation(), true)

	equation = MakeEquation("190", []string{"101", "19"})
	testutils.ExpectToMatchBool(t, equation.HasValidEquation(), false)

	equation = MakeEquation("3267", []string{"81", "40", "27"})
	testutils.ExpectToMatchBool(t, equation.HasValidEquation(), true)
}

func TestHasValidEquationConcat(t *testing.T) {
	// equation := MakeEquation("156", []string{"15", "6"})
	// testutils.ExpectToMatchBool(t, equation.HasValidEquation(), true)

	equation := MakeEquation("7290", []string{"6", "8", "6", "15"})
	testutils.ExpectToMatchBool(t, equation.HasValidEquation(), true)

	// equation = MakeEquation("192", []string{"17", "8", "14"})
	// testutils.ExpectToMatchBool(t, equation.HasValidEquation(), true)

	// equation = MakeEquation("17814", []string{"17", "8", "14"})
	// testutils.ExpectToMatchBool(t, equation.HasValidEquation(), true)
}
