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

func (m *Map) MoveCursorUp() {
	for i := 0; i < len(m.coordinates); i++ {
		fmt.Print("\033[A\033[K") // Move up one line and clear it
	}
}

var purple = "\033[35m" // Purple
var green = "\033[32m"  // Green
var red = "\033[31m"    // Red
var yellow = "\033[33m" // Yellow
var reset = "\033[0m"   // Reset

func (m Map) String() string {
	mapPrinter := "\n"
	for i, row := range m.coordinates {
		for _, coord := range row {
			if coord.Obstacle {
				mapPrinter += (red + "#" + reset)
			} else if coord.Guard != nil {
				mapPrinter += (purple + fmt.Sprint(coord.Guard)) + reset
			} else if coord.Visited {
				mapPrinter += (yellow + getVisitText(coord) + reset)
			} else {
				mapPrinter += (green + "." + reset)
			}
		}
		if i != len(m.coordinates)-1 {
			mapPrinter += "\n"
		}
	}
	return mapPrinter
}

func (m *Map) PrintMapState() {
	m.MoveCursorUp()
	fmt.Print(m)
	// time.Sleep(10 * time.Millisecond)
}

func getVisitText(c *Coordinate) string {
	if c.VisitedType == HORIZONTAL {
		return "\u254C"
	} else if c.VisitedType == VERTICAL {
		return "\u2506"
	} else {
		return "+"
	}
}

type Coordinate struct {
	X           int
	Y           int
	Guard       *Guard
	Obstacle    bool
	Visited     bool
	VisitedType VisitDirection
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

// Direction represents the direction enum
type VisitDirection int

const (
	HORIZONTAL VisitDirection = iota
	VERTICAL
	BOTH
)

func (d VisitDirection) String() string {
	return [...]string{"HORIZONTAL", "VERTICAL", "BOTH"}[d]
}
