package evolution

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestEvolve(t *testing.T) {
	e := Evolution{
		Stones: "0 1 10 99 999",
	}
	testutils.ExpectToMatchString(t, fmt.Sprint(e.Evolve()), "1 2024 1 0 9 9 2021976")
}

func TestEvolveIterationsTest(t *testing.T) {
	iterations := 6
	evolveState := Evolution{
		Stones: "125 17",
	}
	expectedEvolves := []string{
		"253000 1 7",
		"253 0 2024 14168",
		"512072 1 20 24 28676032",
		"512 72 2024 2 0 2 4 2867 6032",
		"1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32",
		"2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2",
	}
	for i := 0; i < iterations; i++ {
		evolveState = evolveState.Evolve()
		fmt.Printf("\n Checking evolution: %v\n", evolveState)
		testutils.ExpectToMatchString(t, fmt.Sprint(evolveState), expectedEvolves[i])
	}
	testutils.ExpectToMatchInt(t, evolveState.CountStones(), 22)
}
