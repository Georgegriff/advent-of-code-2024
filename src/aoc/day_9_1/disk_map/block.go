package diskmap

import (
	"aoc/src/aoc/colors"
	"fmt"
)

type Block struct {
	File *File
}

func (b Block) String() string {
	if b.isEmpty() {
		return colors.PrintColor(".", colors.GRAY)
	}
	return colors.PrintColor(fmt.Sprint(b.File), colors.PURPLE)
}

func MakeFileBlocks(id int, size int) []*Block {
	blocks := make([]*Block, size)
	for i := 0; i < size; i++ {
		file := &File{
			Id: id,
		}
		blocks[i] = &Block{
			File: file,
		}
	}
	return blocks
}

func MakeEmptyBlocks(size int) []*Block {
	blocks := make([]*Block, size)
	for i := 0; i < size; i++ {
		blocks[i] = &Block{}
	}
	return blocks
}

func (b *Block) isEmpty() bool {
	return b.File == nil
}

type File struct {
	Id int
}

func (f File) String() string {
	return colors.PrintColor(fmt.Sprint(f.Id), colors.PURPLE)
}
