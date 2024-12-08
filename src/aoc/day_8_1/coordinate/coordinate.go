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
	if coord.HasAntinode {
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

func (coord *Coordinate) CalculateAntinodePoints(other *Coordinate) []Coordinate {
	multi := 2

	coordinate := coord.findCoordinate(coord.getDistanceVector(other, multi))
	coordinate2 := other.findCoordinate(other.getDistanceVector(coord, multi))

	return []Coordinate{
		coordinate, coordinate2,
	}
}

func (coord *Coordinate) findCoordinate(distance DistanceVector) Coordinate {

	x := coord.X + distance.X
	y := coord.Y + distance.Y

	return Coordinate{X: x, Y: y}
}

func (coord *Coordinate) getDistanceVector(other *Coordinate, multiplier int) DistanceVector {
	// (x2​−x1​,y2​−y1​)
	return DistanceVector{
		X: (other.X - coord.X) * multiplier,
		Y: (other.Y - coord.Y) * multiplier,
	}
}
