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
			file := File{
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
	diskmap := LoadDiskMap("../test-simple.txt")
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0..111....22222"))

}

func TestMoveToNextEmptyBlock(t *testing.T) {
	diskmap := LoadDiskMap("../test-simple.txt")
	update := diskmap.moveToNextEmptyBlock(0, len(diskmap.Blocks)-1)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("02.111....2222."))
	testutils.ExpectToMatchInt(t, update.Start, 1)
	testutils.ExpectToMatchInt(t, update.End, len(diskmap.Blocks)-2)

	update = diskmap.moveToNextEmptyBlock(update.Start, update.End)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("022111....222.."))
	testutils.ExpectToMatchInt(t, update.Start, 2)
	testutils.ExpectToMatchInt(t, update.End, len(diskmap.Blocks)-3)

	update = diskmap.moveToNextEmptyBlock(update.Start, update.End)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0221112...22..."))
	testutils.ExpectToMatchInt(t, update.Start, 6)
	testutils.ExpectToMatchInt(t, update.End, len(diskmap.Blocks)-4)

	update = diskmap.moveToNextEmptyBlock(update.Start, update.End)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("02211122..2...."))
	testutils.ExpectToMatchInt(t, update.Start, 7)
	testutils.ExpectToMatchInt(t, update.End, len(diskmap.Blocks)-5)

	update = diskmap.moveToNextEmptyBlock(update.Start, update.End)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("022111222......"))
	testutils.ExpectToMatchInt(t, update.Start, 8)
	testutils.ExpectToMatchInt(t, update.End, len(diskmap.Blocks)-6)
}

func TestCompactDiskSimple(t *testing.T) {
	diskmap := LoadDiskMap("../test-simple.txt")
	diskmap.CompactDisk(0, len(diskmap.Blocks)-1)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("022111222......"))
}

func TestCompactDisk(t *testing.T) {
	diskmap := LoadDiskMap("../test.txt")
	diskmap.CompactDisk(0, len(diskmap.Blocks)-1)
	testutils.ExpectToMatchString(t, fmt.Sprint(diskmap), formatDiskMapOutput("0099811188827773336446555566.............."))
}

func TestCalculateCheckSum(t *testing.T) {
	diskmap := LoadDiskMap("../test.txt")
	diskmap.CompactDisk(0, len(diskmap.Blocks)-1)
	testutils.ExpectToMatchInt(t, diskmap.CalculateCheckSum(), 1928)
}
