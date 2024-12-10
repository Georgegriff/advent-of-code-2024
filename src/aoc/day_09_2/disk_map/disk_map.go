package diskmap

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"strings"
)

type DiskMap struct {
	Blocks              []*Block
	EmptySpacePositions *OrderedSet[int]
}

func (d DiskMap) String() string {
	rep := ""
	for _, block := range d.Blocks {
		rep += fmt.Sprint(block)
	}
	return rep
}

func (d *DiskMap) hasFreeSpaceAt(startPosition, size int) bool {
	for i := 0; i < size; i++ {
		space := d.EmptySpacePositions.Contains(startPosition + i)
		if !space {
			return false
		}
	}

	return true
}

func (d *DiskMap) getEmptySpace(n int) (int, error) {
	el, err := d.EmptySpacePositions.Nth(n)

	if err != nil {
		return -1, err
	}
	return el, nil
}

func (d *DiskMap) getFirstEmptySpace(n int) (int, error) {
	el, err := d.EmptySpacePositions.Nth(n)

	if err != nil {
		return -1, err
	}
	return el, nil
}

func (d *DiskMap) moveFileToEmptySpace(emptySpaceOffset int, fileStartPosition int, file *BlockFile) error {
	firstEmptySpace, err := d.getEmptySpace(emptySpaceOffset)
	if err != nil {
		return fmt.Errorf("no empty left")
	}
	if firstEmptySpace > fileStartPosition {
		return fmt.Errorf("cannot move file: %d", file.Id)
	}
	if d.hasFreeSpaceAt(firstEmptySpace, file.Size) {
		for i := 0; i < file.Size; i++ {
			d.Blocks[fileStartPosition-i].File = nil
			d.Blocks[firstEmptySpace+i].File = file
			d.EmptySpacePositions.Remove(firstEmptySpace + i)
		}
		return nil
	}
	emptySpaceOffset += 1
	return d.moveFileToEmptySpace(emptySpaceOffset, fileStartPosition, file)

}

func (d *DiskMap) moveFileIfFits(startPosition int, endPosition int) (*BlockUpdate, error) {
	blocks := d.Blocks
	for startPosition < endPosition {
		endPositionBlock := blocks[endPosition]
		if endPositionBlock.isEmpty() {
			endPosition--
			continue
		}
		fileToMove := blocks[endPosition].File

		d.moveFileToEmptySpace(0, endPosition, fileToMove)
		endPosition = endPosition - fileToMove.Size
		break
	}
	return &BlockUpdate{
		Start: startPosition,
		End:   endPosition,
	}, nil
}

func (d *DiskMap) CompactDisk(startPosition int, endPosition int) {

	if endPosition <= startPosition {
		return
	}
	update, _ := d.moveFileIfFits(startPosition, endPosition)

	d.CompactDisk(update.Start, update.End)
}

func (d *DiskMap) CalculateCheckSum() int {
	sum := 0
	for i, block := range d.Blocks {
		if block.isEmpty() {
			continue
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
	orderedSet := NewOrderedSet[int]()
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
					spaces := MakeEmptyBlocks(spaceSize, len(blocks), orderedSet)
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
		Blocks:              blocks,
		EmptySpacePositions: orderedSet,
	}
}

func LoadOutputFormat(line string) *DiskMap {
	blocks := []*Block{}

	numsStr := strings.Split(line, "")
	// process file and block space at once
	lastNum := -1
	var file *BlockFile
	orderedSet := NewOrderedSet[int]()
	for i := 0; i < len(numsStr); i++ {
		str := numsStr[i]
		if str == "." {
			orderedSet.Add(i)
			blocks = append(blocks, &Block{})
		} else {
			num := utils.ToInt(str)
			if lastNum != num {
				lastNum = num
				file = &BlockFile{
					Id:   lastNum,
					Size: 1,
				}
			} else {
				file.Size++
			}
			blocks = append(blocks, &Block{
				File: file,
			})
		}

	}

	return &DiskMap{
		Blocks:              blocks,
		EmptySpacePositions: orderedSet,
	}
}

type BlockUpdate struct {
	Start int
	End   int
}
