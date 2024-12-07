package main

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
	"strings"
)

type Map struct {
	coordinates [][]*Coordinate
	Guard       *Guard
}

func (m Map) String() string {
	mapPrinter := ""
	for i, row := range m.coordinates {
		for _, coord := range row {
			if coord.Obstacle {
				mapPrinter += "#"
			} else if coord.Guard != nil {
				mapPrinter += fmt.Sprint(coord.Guard)
			} else if coord.Visited {
				mapPrinter += "X"
			} else {
				mapPrinter += "."
			}
		}
		if i != len(m.coordinates)-1 {
			mapPrinter += "\n"
		}
	}
	return mapPrinter
}

type Coordinate struct {
	X        int
	Y        int
	Guard    *Guard
	Obstacle bool
	Visited  bool
}

func NewCoordinate(value string, x int, y int) *Coordinate {
	coordinate := Coordinate{
		X: x,
		Y: y,
	}

	if value == "#" || value == "." {
		if value == "#" {
			coordinate.Obstacle = true
		}
		return &coordinate
	}
	guardPosition, err := GetGuardDirection(value)
	if err != nil {
		log.Fatal(err)
	}
	coordinate.Guard = NewGuard(&coordinate, guardPosition)
	return &coordinate
}

func LoadMap(path string) *Map {
	file := readfile.Open(path)
	defer file.Close()
	rowCounter := 0
	coordinates := [][]*Coordinate{}
	var guard *Guard
	err := readfile.ReadLine(file, func(line string) error {
		positionOnLine := strings.Split(line, "")

		row := []*Coordinate{}
		for x, position := range positionOnLine {
			coord := NewCoordinate(position, x, rowCounter)
			if coord.Guard != nil {
				guard = coord.Guard
			}
			row = append(row, coord)
		}
		coordinates = append(coordinates, row)
		rowCounter++
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Map{
		coordinates: coordinates,
		Guard:       guard,
	}
}
