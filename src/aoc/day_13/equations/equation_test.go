package equations

import (
	"aoc/src/testutils"
	"testing"
)

func TestSolveEquation(t *testing.T) {
	equation1 := Equation{
		output:       8400,
		aCoefficient: 94,
		bCoefficient: 22,
	}
	equation2 := Equation{
		output:       5400,
		aCoefficient: 34,
		bCoefficient: 67,
	}
	testutils.ExpectToMatchBool(t, equation1.SolveEquations(&equation2), true)
	testutils.ExpectToMatchInt(t, equation1.a, 80)
	testutils.ExpectToMatchInt(t, equation1.b, 40)
	testutils.ExpectToMatchInt(t, equation2.a, 80)
	testutils.ExpectToMatchInt(t, equation2.b, 40)

	testutils.ExpectToMatchInt(t, equation1.GetCost(3, 1), 280)
}
