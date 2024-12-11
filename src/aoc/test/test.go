package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	initial := []string{"70949", "6183", "4", "3825336", "613971", "0", "15", "182"}
	blinks := 75

	stoneCount := simulateEfficient(initial, blinks)
	totalStones := 0
	for _, count := range stoneCount {
		totalStones += count
	}
	fmt.Println("Total stones after", blinks, "blinks:", totalStones)
}

// simulateEfficient efficiently simulates stone evolution
func simulateEfficient(initial []string, blinks int) map[string]int {
	// Initialize map to count stones
	stoneMap := make(map[string]int)
	for _, stone := range initial {
		stoneMap[stone]++
	}

	for b := 0; b < blinks; b++ {
		fmt.Printf("\n blink count: %d \n", b)
		newStoneMap := make(map[string]int)
		for stone, count := range stoneMap {
			for _, result := range evolveStone(stone) {
				newStoneMap[result] += count
			}
		}
		stoneMap = newStoneMap
	}

	return stoneMap
}

// evolveStone applies the transformation rules to a single stone
func evolveStone(stone string) []string {
	if stone == "0" {
		return []string{"1"}
	}

	if len(stone)%2 == 0 {
		mid := len(stone) / 2
		left := strings.TrimLeft(stone[:mid], "0")
		right := strings.TrimLeft(stone[mid:], "0")
		if left == "" {
			left = "0"
		}
		if right == "" {
			right = "0"
		}
		return []string{left, right}
	}

	// Multiply by 2024
	val := new(big.Int)
	val.SetString(stone, 10)
	val.Mul(val, big.NewInt(2024))
	return []string{val.String()}
}
