package maze

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestLoadMaze(t *testing.T) {
	maze := LoadMaze("../test.txt", EAST)

	testutils.ExpectToMatchString(t, fmt.Sprint(maze), `
###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`)
}

func TestBuildGraph(t *testing.T) {
	maze := LoadMaze("../test.txt", EAST)
	score := maze.GetBestPathScore()
	testutils.ExpectToMatchInt(t, score, 7036)

}
