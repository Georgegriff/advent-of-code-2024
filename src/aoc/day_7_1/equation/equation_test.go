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
