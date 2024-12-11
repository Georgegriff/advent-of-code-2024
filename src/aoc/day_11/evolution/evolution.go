package evolution

import (
	"aoc/src/aoc/utils"
	"fmt"
	"strings"
)

type StoneMap = map[string]int
type Evolution struct {
	Stones   string
	StoneMap StoneMap
}

func MakeEvolution(Stones string, stoneMap StoneMap) *Evolution {
	if stoneMap == nil {
		stoneMap = make(StoneMap)
		stones := strings.Split(Stones, " ")
		for _, stone := range stones {
			stoneMap[stone]++
		}
	}
	return &Evolution{
		Stones:   Stones,
		StoneMap: stoneMap,
	}
}

func (e Evolution) String() string {
	return e.Stones
}

func (e *Evolution) CountStones() int {
	sum := 0
	for _, v := range e.StoneMap {
		sum += v
	}

	return sum
}

func (e *Evolution) Evolve(iterations int) *Evolution {

	stoneMap := e.evolveStones(iterations)
	stones := make([]string, len(stoneMap))
	i := 0
	for k := range stoneMap {
		stones[i] = k
		i++
	}
	return MakeEvolution(
		strings.Join(stones, " "),
		stoneMap,
	)
}

func (e *Evolution) evolveStones(iterations int) StoneMap {
	stoneMap := e.StoneMap
	var evolution []string
	for i := 0; i < iterations; i++ {
		newIterationMap := make(StoneMap)
		for stone, count := range stoneMap {
			evolution = evolveStone(stone)
			for _, result := range evolution {
				newIterationMap[result] += count
			}
		}
		stoneMap = newIterationMap
	}

	return stoneMap
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
