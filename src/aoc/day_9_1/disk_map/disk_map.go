package diskmap

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"strings"
)

type DiskMap struct {
	Blocks []*Block
}

func (d DiskMap) String() string {
	rep := ""
	for _, block := range d.Blocks {
		rep += fmt.Sprint(block)
	}
	return rep
}

func (d *DiskMap) moveToNextEmptyBlock(startPosition, endPosition int) BlockUpdate {
	blocks := d.Blocks
	for startPosition < endPosition {
		endPositionBlock := blocks[endPosition]
		if endPositionBlock.isEmpty() {
			endPosition--
			continue
		}
		startBlock := blocks[startPosition]
		if startBlock.isEmpty() {
			fileToMove := blocks[endPosition].File
			blocks[endPosition].File = nil
			blocks[startPosition].File = fileToMove
			endPosition--
			break
		} else {
			startPosition++
		}
	}
	return BlockUpdate{
		Start: startPosition,
		End:   endPosition,
	}
}

func (d *DiskMap) CompactDisk(startPosition, endPosition int) {
	if endPosition <= startPosition {
		return
	}
	update := d.moveToNextEmptyBlock(startPosition, endPosition)
	d.CompactDisk(update.Start, update.End)
}

func (d *DiskMap) CalculateCheckSum() int {
	sum := 0
	for i, block := range d.Blocks {
		if block.isEmpty() {
			break
		}
		sum += i * block.File.Id
	}
	return sum
}

func LoadDiskMap(path string) *DiskMap {
	file := readfile.Open(path)
	defer file.Close()
	blockIdCounter := 0
	blocks := []*Block{}
	err := readfile.ReadLine(file, func(line string) error {
		numsStr := strings.Split(line, "")
		// process file and block space at once
		for i := 0; i < len(numsStr); i += 2 {
			fileSize := utils.ToInt(numsStr[i])

			if fileSize > 0 {
				files := MakeFileBlocks(blockIdCounter, fileSize)
				blocks = append(blocks, files...)
			}
			// check if at end of the array
			if i+1 < len(numsStr) {
				spaceSize := utils.ToInt(numsStr[i+1])
				if spaceSize > 0 {
					spaces := MakeEmptyBlocks(spaceSize)
					blocks = append(blocks, spaces...)
				}
				blockIdCounter++
			}

		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &DiskMap{
		Blocks: blocks,
	}
}

type BlockUpdate struct {
	Start int
	End   int
}
