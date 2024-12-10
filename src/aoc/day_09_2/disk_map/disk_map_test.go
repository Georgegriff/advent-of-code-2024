package diskmap

import (
	"aoc/src/aoc/utils"
	"aoc/src/testutils"
	"fmt"
	"strings"
	"testing"
)

func formatDiskMapOutput(input string) string {
	entries := strings.Split(input, "")
	printer := ""
	for _, entry := range entries {
		if entry == "." {
			printer += fmt.Sprint(Block{})
		} else {
			num := utils.ToInt(entry)
			file := BlockFile{
				Id: num,
			}
			printer += fmt.Sprint(Block{
				File: &file,
			})
		}
	}
	return printer
}

func TestLoadDiskMap(t *testing.T) {
	diskmap := LoadDiskMap("../test.txt")
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("00...111...2...333.44.5555.6666.777.888899"))
}

func TestMoveFileIfFits(t *testing.T) {
	diskmap := LoadDiskMap("../test.txt")
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("00...111...2...333.44.5555.6666.777.888899"))

	update, err := diskmap.moveFileIfFits(0, len(diskmap.Blocks)-1)
	testutils.ExpectToErrorNil(t, err)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.111...2...333.44.5555.6666.777.8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// 8888 can't move
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.111...2...333.44.5555.6666.777.8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.1117772...333.44.5555.6666.....8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// 6666 can't move
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.1117772...333.44.5555.6666.....8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// 5555 can't move
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.1117772...333.44.5555.6666.....8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// 44 moved
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.111777244.333....5555.6666.....8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// 333 can't move
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.111777244.333....5555.6666.....8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// 44 moved
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099.111777244.333....5555.6666.....8888.."))

	update, err = diskmap.moveFileIfFits(update.Start, update.End)
	testutils.ExpectToErrorNil(t, err)
	// Move 2
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("00992111777.44.333....5555.6666.....8888.."))

}

func TestCompactDisk(t *testing.T) {
	diskmap := LoadDiskMap("../test.txt")
	diskmap.CompactDisk(0, len(diskmap.Blocks)-1)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("00992111777.44.333....5555.6666.....8888.."))
}

func TestCalculateCheckSumOrdered(t *testing.T) {
	diskmap := LoadOutputFormat("00992111777.44.333....5555.6666.....8888..")
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("00992111777.44.333....5555.6666.....8888.."))

	testutils.ExpectToMatchInt(t, diskmap.CalculateCheckSum(), 2858)

}

// func TestCalculateCheckSum(t *testing.T) {
// 	diskmap := LoadDiskMap("../test.txt")
// 	diskmap.CompactDisk(0, len(diskmap.Blocks)-1)
// 	testutils.ExpectToMatchInt(t, diskmap.CalculateCheckSum(), 2858)
// }
