package evolution

import (
	"aoc/src/testutils"
	"testing"
)

func TestEvolve(t *testing.T) {
	e := MakeEvolution("0 1 10 99 999", nil).Evolve(1)

	testutils.ExpectToMatchInt(t, e.StoneMap["1"], 2)
	testutils.ExpectToMatchInt(t, e.StoneMap["2024"], 1)
	testutils.ExpectToMatchInt(t, e.StoneMap["0"], 1)
	testutils.ExpectToMatchInt(t, e.StoneMap["9"], 2)
	testutils.ExpectToMatchInt(t, e.StoneMap["2021976"], 1)

}

func TestEvolveIterationsTest(t *testing.T) {
	iterations := 6
	evolveState := MakeEvolution("125 17", nil)

	evolveState = evolveState.Evolve(iterations)

	testutils.ExpectToMatchInt(t, evolveState.CountStones(), 22)
}
