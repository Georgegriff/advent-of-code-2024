package topmap

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestLoadTopMapSimple(t *testing.T) {
	topMap := LoadMap("../simple.txt")
	fmt.Print(topMap)
	testutils.ExpectToMatchString(t, fmt.Sprint(topMap),
		`0123
1234
8765
9876`)

	trailHeads := topMap.TrailHeads
	testutils.ExpectToMatchInt(t, len(trailHeads), 1)
}

func TestCalculateTrailHeadPathsSimple(t *testing.T) {
	topMap := LoadMap("../simple.txt")

	paths := topMap.GetTrailHeadPaths(topMap.TrailHeads[0])
	seenPaths := make(map[string]bool)
	for _, path := range paths {
		pathStr := PrintPath(topMap, path)
		// check unique
		seenPaths[pathStr] = true
		// fmt.Printf("\nchecking path: %v", pathStr)
		testutils.ExpectToMatchBool(t, ValidateTrail(topMap, path), true)
	}
}

func TestCalculateTrailHeadPaths(t *testing.T) {
	topMap := LoadMap("../test.txt")

	paths := topMap.GetTrailHeadPaths(topMap.TrailHeads[1])
	seenPaths := make(map[string]bool)
	for _, path := range paths {
		pathStr := PrintPath(topMap, path)
		// check unique
		seenPaths[pathStr] = true
		testutils.ExpectToMatchBool(t, ValidateTrail(topMap, path), true)
	}
}
