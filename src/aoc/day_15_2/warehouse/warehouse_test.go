package warehouse

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestLoadWareHouse(t *testing.T) {

	warehouse := LoadWarehouse("../simple.txt", 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
################
##....[]..[]..##
####@...[]....##
##......[]....##
##..##..[]....##
##......[]....##
##............##
################`)

}

func TestLoadWareHouseSimple2(t *testing.T) {

	warehouse := LoadWarehouse("../simple-2.txt", 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############`)

}

func TestLoadWareHouseLarger(t *testing.T) {

	warehouse := LoadWarehouse("../test.txt", 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
####################
##....[]....[]..[]##
##............[]..##
##..[][]....[]..[]##
##....[]@.....[]..##
##[]##....[]......##
##[]....[]....[]..##
##..[][]..[]..[][]##
##........[]......##
####################`)

}

func TestMoveBoxes(t *testing.T) {
	warehouse := LoadWarehouse("../simple-2.txt", 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##......##..##
##..........##
##....[][]@.##
##....[]....##
##..........##
##############`)
	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##......##..##
##..........##
##...[][]@..##
##....[]....##
##..........##
##############`)

}

func TestEdgeCaseMove(t *testing.T) {
	warehouse := LoadWarehouse("../edge-case-2.txt", 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##..........##
##..[]##....##
##....[]@...##
##..........##
##############`)
	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##..........##
##..[]##....##
##...[]@....##
##..........##
##############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##..........##
##..[]##....##
##...[].....##
##.....@....##
##############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##..........##
##..[]##....##
##...[].....##
##....@.....##
##############`)
	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##..........##
##..[]##....##
##...[].....##
##...@......##
##############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
##############
##..........##
##..[]##....##
##...[].....##
##...@......##
##############`)
}

func TestEdgeCase4Move(t *testing.T) {
	warehouse := LoadWarehouse("../edge-case-4.txt", 2)
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##........##
##..[]....##
##..[][]@.##
##..[]....##
##........##
############`)
	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##........##
##..[]....##
##.[][]@..##
##..[]....##
##........##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##........##
##..[]....##
##.[][]...##
##..[].@..##
##........##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##........##
##..[]....##
##.[][]...##
##..[]....##
##.....@..##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##........##
##..[]....##
##.[][]...##
##..[]....##
##....@...##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##........##
##..[]....##
##.[][]...##
##..[]....##
##...@....##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##....##..##
##..[]....##
##.[][]...##
##..[]....##
##...@....##
##........##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##..[]##..##
##.[][]...##
##..[]....##
##...@....##
##........##
##........##
############`)

	warehouse.Next()
	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
############
##........##
##..[]##..##
##.[][]...##
##..[]....##
##...@....##
##........##
##........##
############`)

}

// func TestPushMultipleUp(t *testing.T) {
// 	warehouse := LoadWarehouse("../simple-2.txt", 2)
// 	warehouse.Next()
// 	warehouse.Next()
// 	warehouse.Next()
// 	warehouse.Next()
// 	warehouse.Next()

// 	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
// ##############
// ##......##..##
// ##..........##
// ##...[][]...##
// ##....[]....##
// ##.....@....##
// ##############`)
// 	warehouse.Next()
// 	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
// ##############
// ##......##..##
// ##...[][]...##
// ##....[]....##
// ##.....@....##
// ##..........##
// ##############`)
// }

// func TestProcessSimple(t *testing.T) {
// 	warehouse := LoadWarehouse("../simple-2.txt", 2)
// 	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
// ##############
// ##......##..##
// ##..........##
// ##....[][]@.##
// ##....[]....##
// ##..........##
// ##############`)
// 	warehouse.Process()
// 	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
// ##############
// ##...[].##..##
// ##...@.[]...##
// ##....[]....##
// ##..........##
// ##..........##
// ##############`)

// }

// func TestProcess(t *testing.T) {
// 	warehouse := LoadWarehouse("../test.txt", 2)

// 	warehouse.Process()
// 	testutils.ExpectToMatchString(t, fmt.Sprint(warehouse), `
// ####################
// ##[].......[].[][]##
// ##[]...........[].##
// ##[]........[][][]##
// ##[]......[]....[]##
// ##..##......[]....##
// ##..[]............##
// ##..@......[].[][]##
// ##......[][]..[]..##
// ####################`)

// }
