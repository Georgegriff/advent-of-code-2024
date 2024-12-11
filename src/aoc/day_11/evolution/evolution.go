package evolution

import (
	"aoc/src/aoc/utils"
	"fmt"
	"strings"
)

type Evolution struct {
	Stones string
}

func (e Evolution) String() string {
	return e.Stones
}

func (e *Evolution) CountStones() int {
	return len(strings.Split(e.Stones, " "))
}

func (e *Evolution) Evolve() Evolution {

	return Evolution{
		Stones: evolveStones(e.Stones),
	}
}

func evolveStones(stonesInput string) string {
	stones := strings.Split(stonesInput, " ")
	stoneBuilder := []string{}
	for _, stone := range stones {
		stoneBuilder = append(stoneBuilder, evolveStone(stone)...)
	}

	return strings.Join(stoneBuilder, " ")
}

func evolveStone(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	} else if len(stone)%2 == 0 {
		stoneOne, stoneTwo := splitStringInHalf(stone)

		return []string{trimLeadingZeros(stoneOne), trimLeadingZeros(stoneTwo)}
	} else {
		num := utils.ToInt64(stone)
		updatedNum := num * 2024
		return []string{fmt.Sprintf("%d", updatedNum)}
	}
}

func splitStringInHalf(s string) (string, string) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}

func trimLeadingZeros(input string) string {
	output := strings.TrimLeft(input, "0")
	if output == "" {
		return "0"
	}
	return output
}
