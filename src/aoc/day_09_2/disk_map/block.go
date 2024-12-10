package diskmap

import (
	"aoc/src/aoc/colors"
	"fmt"
)

type Block struct {
	File *BlockFile
}

func (b Block) String() string {
	if b.isEmpty() {
		return colors.PrintColor(".", colors.GRAY)
	}
	return colors.PrintColor(fmt.Sprint(b.File), colors.PURPLE)
}

func MakeFileBlocks(id int, size int) []*Block {
	blocks := make([]*Block, size)
	file := &BlockFile{
		Id:   id,
		Size: size,
	}
	for i := 0; i < size; i++ {
		blocks[i] = &Block{
			File: file,
		}
	}
	return blocks
}

func MakeEmptyBlocks(size int, spaceOffset int, om *OrderedSet[int]) []*Block {
	blocks := make([]*Block, size)
	for i := 0; i < size; i++ {
		blocks[i] = &Block{}
		om.Add(spaceOffset + i)
	}
	return blocks
}

func (b *Block) isEmpty() bool {
	return b.File == nil
}

type BlockFile struct {
	Id   int
	Size int
}

func (f BlockFile) String() string {
	return colors.PrintColor(fmt.Sprint(f.Id), colors.PURPLE)
}
