package coordinate

import (
	"fmt"
)

var purple = "\033[35m" // Purple
var green = "\033[32m"  // Green
var red = "\033[31m"    // Red
var yellow = "\033[33m" // Yellow
var reset = "\033[0m"   // Reset

func NewCoordinate(value string, x int, y int) *Coordinate {
	return &Coordinate{
		X:           x,
		Y:           y,
		AntennaType: AntennaType(value),
	}
}

type DistanceVector struct {
	X int
	Y int
}

const EMPTY_POINT = "."
const ANTINODE = "#"

type AntennaType string

type Coordinate struct {
	X           int
	Y           int
	AntennaType AntennaType
	HasAntinode bool
}

func (coord Coordinate) String() string {
	if coord.HasAntinode && coord.AntennaType != "" {
		return (yellow + fmt.Sprint(coord.AntennaType)) + reset
	} else if coord.HasAntinode {
		return (red + ANTINODE + reset)
	} else if coord.AntennaType != "" {
		return (purple + fmt.Sprint(coord.AntennaType)) + reset
	} else {
		return EMPTY_POINT
	}
}

func (c Coordinate) GoString() string {
	antennaStr := fmt.Sprint(c)
	return fmt.Sprintf("\n%s[X:%d,Y:%d]", antennaStr, c.X, c.Y)
}

func (coord *Coordinate) HasAntenna() bool {
	return string(coord.AntennaType) != ""
}

func (coord *Coordinate) CalculateAntinodePoints(other *Coordinate, yMax int, xMax int) []Coordinate {
	var points []Coordinate

	distance := coord.getDistanceVector(other, 1)

	// start at negativeX to push the start backwards, and loop until max of grid
	for i := -coord.X; i <= xMax-coord.X; i++ {
		x := coord.X + distance.X*i
		y := coord.Y + distance.Y*i
		points = append(points, Coordinate{X: x, Y: y})
	}
	return points
}

func (coord *Coordinate) getDistanceVector(other *Coordinate, multiplier int) DistanceVector {
	// (x2​−x1​,y2​−y1​)
	return DistanceVector{
		X: (other.X - coord.X) * multiplier,
		Y: (other.Y - coord.Y) * multiplier,
	}
}
