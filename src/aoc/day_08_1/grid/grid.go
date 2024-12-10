package grid

import (
	"aoc/src/aoc/day_08_1/coordinate"
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
	"strings"
)

type Antennas map[coordinate.AntennaType][]*coordinate.Coordinate

type Grid struct {
	Coordinates [][]*coordinate.Coordinate
	Antennas    Antennas
	Antinodes   map[*coordinate.Coordinate]bool
}

func (g Grid) String() string {
	mapPrinter := "\n"
	for i, row := range g.Coordinates {
		for _, coord := range row {
			mapPrinter += fmt.Sprint(coord)
		}
		if i != len(g.Coordinates)-1 {
			mapPrinter += "\n"
		}
	}
	return mapPrinter
}

func (g Grid) AddAntinodeToCoordinates(coordinates []coordinate.Coordinate) {
	YMax := len(g.Coordinates) - 1
	for _, coordinate := range coordinates {
		if coordinate.Y < 0 || coordinate.Y > YMax {
			continue
		}
		row := g.Coordinates[coordinate.Y]
		xMax := len(row) - 1
		if coordinate.X < 0 || coordinate.X > xMax {
			continue
		}
		realCoordinate := g.Coordinates[coordinate.Y][coordinate.X]
		realCoordinate.HasAntinode = true
		g.Antinodes[realCoordinate] = true
	}
}

func LoadGrid(path string) *Grid {
	file := readfile.Open(path)
	defer file.Close()
	rowCounter := 0
	var antennas Antennas = make(Antennas)
	coordinates := [][]*coordinate.Coordinate{}
	grid := Grid{
		Coordinates: coordinates,
		Antennas:    antennas,
		Antinodes:   make(map[*coordinate.Coordinate]bool),
	}
	err := readfile.ReadLine(file, func(line string) error {

		points := strings.Split(line, "")
		row := []*coordinate.Coordinate{}
		for i, point := range points {
			var coord *coordinate.Coordinate
			if point == coordinate.EMPTY_POINT {
				coord = coordinate.NewCoordinate("", i, rowCounter)
			} else {
				coord = coordinate.NewCoordinate(point, i, rowCounter)
				if coord.HasAntenna() && grid.Antennas[coord.AntennaType] == nil {
					grid.Antennas[coord.AntennaType] = []*coordinate.Coordinate{}
				}
				antennaForType := grid.Antennas[coord.AntennaType]
				grid.Antennas[coord.AntennaType] = append(antennaForType, coord)
			}
			row = append(row, coord)
		}
		grid.Coordinates = append(grid.Coordinates, row)
		rowCounter++
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &grid

}
