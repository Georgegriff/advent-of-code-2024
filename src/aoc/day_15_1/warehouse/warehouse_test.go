package warehouse

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestLoadWareHouse(t *testing.T) {

	warehouse := LoadWarehouse("../simple.txt")
	testutils.ExpectToMatchInt(t, len(warehouse.boxes), 6)
	testutils.ExpectToMatchInt(t, warehouse.robot.X, 2)
	testutils.ExpectToMatchInt(t, warehouse.robot.Y, 2)
	testutils.ExpectToMatchString(t, warehouse.robot.Instructions, "<^^>>>vv<v>>v<<")
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)

}

func TestMoveRobot(t *testing.T) {
	warehouse := LoadWarehouse("../simple.txt")
	testutils.ExpectToMatchString(t, warehouse.robot.Instructions, "<^^>>>vv<v>>v<<")
	testutils.ExpectToMatchInt(t, warehouse.robot.CurrentInstruction, 0)

	// doesn't move because wall
	warehouse.Next()
	testutils.ExpectToMatchInt(t, warehouse.robot.CurrentInstruction, 1)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
	// moves
	warehouse.Next()
	testutils.ExpectToMatchInt(t, warehouse.robot.CurrentInstruction, 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)

	warehouse.Next()
	// wall
	testutils.ExpectToMatchInt(t, warehouse.robot.CurrentInstruction, 3)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#.@O.O.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)

	warehouse.Next()
	testutils.ExpectToMatchInt(t, warehouse.robot.CurrentInstruction, 4)
	// move box
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)

}

func TestMoveBoxes(t *testing.T) {
	warehouse := LoadWarehouse("../push-box.txt")
	testutils.ExpectToMatchInt(t, warehouse.robot.CurrentInstruction, 0)
	// move box
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#..@OO.#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)

	warehouse.Next()
	// move multiple boxes
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#...@OO#
##..O..#
#...O..#
#.#.O..#
#...O..#
#......#
########`)
}

func TestMovement(t *testing.T) {
	warehouse := LoadWarehouse("../simple.txt")

	warehouse.Process()
	// move multiple boxes
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
########
#....OO#
##.....#
#.....O#
#.#O@..#
#...O..#
#...O..#
########`)
}

func TestMovementLarger(t *testing.T) {
	warehouse := LoadWarehouse("../test.txt")

	warehouse.Process()
	// move multiple boxes
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`)
}
