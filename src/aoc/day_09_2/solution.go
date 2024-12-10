package main

import (
	diskmap "aoc/src/aoc/day_09_2/disk_map"
	"fmt"
)

func Solve(
	path string,
) int {

	diskmap := diskmap.LoadDiskMap(path)

	diskmap.CompactDisk(0, len(diskmap.Blocks)-1)

	return diskmap.CalculateCheckSum()
}

func main() {
	fmt.Printf("\nThe answer is %#v\n", Solve("./input.txt"))
}
