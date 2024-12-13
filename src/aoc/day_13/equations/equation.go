package equations

import "log"

type Equation struct {
	output       int
	a            int
	aCoefficient int
	b            int
	bCoefficient int
}

func MakeEquation(
	output, aCoefficient, bCoefficient, adjustmentOutput int,
) *Equation {
	return &Equation{
		output:       output + adjustmentOutput,
		aCoefficient: aCoefficient,
		bCoefficient: bCoefficient,
	}
}

func (equation *Equation) SolveEquations(equation2 *Equation) bool {

	eq1B := Equation{
		output:       equation.output * equation2.bCoefficient,
		aCoefficient: equation.aCoefficient * equation2.bCoefficient,
		bCoefficient: equation.bCoefficient * equation2.bCoefficient,
	}

	eq2B := Equation{
		output:       equation2.output * equation.bCoefficient,
		aCoefficient: equation2.aCoefficient * equation.bCoefficient,
		bCoefficient: equation2.bCoefficient * equation.bCoefficient,
	}

	if eq1B.bCoefficient != eq2B.bCoefficient {
		log.Fatal("could not eliminate be")
	}

	a := (eq1B.output - eq2B.output) / (eq1B.aCoefficient - eq2B.aCoefficient)

	equation.a = a
	equation2.a = a
	b := (equation.output - equation.a*equation.aCoefficient) / equation.bCoefficient
	equation.b = b
	equation2.b = b

	return (equation.a*equation.aCoefficient+equation.b*equation.bCoefficient == equation.output) &&
		(equation2.a*equation2.aCoefficient+equation2.b*equation2.bCoefficient == equation2.output)

}

func (equation *Equation) GetCost(aCost, bCost int) int {
	return aCost*equation.a + bCost*equation.b
}
