package coordinate

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestGetDistanceVector(t *testing.T) {

	one := NewCoordinate("A", 7, 7)
	two := NewCoordinate("B", 8, 8)

	distanceOne := one.getDistanceVector(two, 1)
	testutils.ExpectToMatchInt(t, distanceOne.X, 1)
	testutils.ExpectToMatchInt(t, distanceOne.Y, 1)

	distanceTwo := two.getDistanceVector(one, 1)
	testutils.ExpectToMatchInt(t, distanceTwo.X, -1)
	testutils.ExpectToMatchInt(t, distanceTwo.Y, -1)

	distanceOne = one.getDistanceVector(two, 2)
	testutils.ExpectToMatchInt(t, distanceOne.X, 2)
	testutils.ExpectToMatchInt(t, distanceOne.Y, 2)

	distanceTwo = two.getDistanceVector(one, 2)
	testutils.ExpectToMatchInt(t, distanceTwo.X, -2)
	testutils.ExpectToMatchInt(t, distanceTwo.Y, -2)
}

func TestCalculateAntinodePoints(t *testing.T) {
	xMax := 11
	yMax := 11

	one := NewCoordinate("A", 8, 8)
	two := NewCoordinate("B", 9, 9)

	points := one.CalculateAntinodePoints(two, xMax, yMax)
	testutils.ExpectToMatchInt(t, len(points), 12)

	// this line should be drawn from top left diagonal to the other for bcased on the input coordinations
	for i := 0; i < len(points); i++ {
		point := points[i]
		testutils.ExpectToMatchString(t, fmt.Sprintf("[%d, %d]", point.X, point.Y), fmt.Sprintf("[%d, %d]", i, i))
	}

}
