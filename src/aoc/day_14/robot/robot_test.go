package robot

import (
	"aoc/src/testutils"
	"testing"
)

func TestMoveRobot(t *testing.T) {

	boundsX := 11
	boundsY := 7

	robot := &Robot{
		X: 2,
		Y: 4,
		Velocity: Velocity{
			X: 2,
			Y: -3,
		},
	}

	updatedRobot := robot.CalculatePositionIn(1, boundsX, boundsY)
	testutils.ExpectToMatchInt(t, updatedRobot.X, 4)
	testutils.ExpectToMatchInt(t, updatedRobot.Y, 1)

	updatedRobot = robot.CalculatePositionIn(2, boundsX, boundsY)
	testutils.ExpectToMatchInt(t, updatedRobot.X, 6)
	testutils.ExpectToMatchInt(t, updatedRobot.Y, 5)

	updatedRobot = robot.CalculatePositionIn(3, boundsX, boundsY)
	testutils.ExpectToMatchInt(t, updatedRobot.X, 8)
	testutils.ExpectToMatchInt(t, updatedRobot.Y, 2)

	updatedRobot = robot.CalculatePositionIn(4, boundsX, boundsY)
	testutils.ExpectToMatchInt(t, updatedRobot.X, 10)
	testutils.ExpectToMatchInt(t, updatedRobot.Y, 6)

	updatedRobot = robot.CalculatePositionIn(5, boundsX, boundsY)
	testutils.ExpectToMatchInt(t, updatedRobot.X, 1)
	testutils.ExpectToMatchInt(t, updatedRobot.Y, 3)
}
