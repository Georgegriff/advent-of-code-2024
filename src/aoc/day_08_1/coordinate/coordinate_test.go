package coordinate

import (
	"aoc/src/testutils"
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

	one := NewCoordinate("A", 8, 8)
	two := NewCoordinate("B", 9, 9)

	points := one.CalculateAntinodePoints(two)
	testutils.ExpectToMatchInt(t, len(points), 2)

	antiNodeOne := points[0]
	antiNodeTwo := points[1]

	testutils.ExpectToMatchInt(t, antiNodeOne.X, 10)
	testutils.ExpectToMatchInt(t, antiNodeOne.Y, 10)

	testutils.ExpectToMatchInt(t, antiNodeTwo.X, 7)
	testutils.ExpectToMatchInt(t, antiNodeTwo.Y, 7)

	one = NewCoordinate("A", 0, 0)
	two = NewCoordinate("B", 1, 0)
	points = one.CalculateAntinodePoints(two)

	testutils.ExpectToMatchInt(t, points[0].X, 2)
	testutils.ExpectToMatchInt(t, points[0].Y, 0)

	testutils.ExpectToMatchInt(t, points[1].X, -1)
	testutils.ExpectToMatchInt(t, points[1].Y, 0)
}
