package robot

import (
	"aoc/src/testutils"
	"fmt"
	"testing"
)

func TestLoadLobby(t *testing.T) {

	lobby := LoadLobby("../test.txt", 11, 7)
	initialLobby := `
1.12.......
...........
...........
......11.11
1.1........
.........1.
.......1...`
	testutils.ExpectToMatchString(t, fmt.Sprint(lobby), initialLobby)
	testutils.ExpectToMatchInt(t, len(lobby.RobotsInQuadrants), 4)

	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[0], 4)
	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[1], 0)
	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[2], 2)
	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[3], 2)

}

func TestMoveRobots(t *testing.T) {

	lobby := LoadLobby("../test.txt", 11, 7)
	after100 := `
......2..1.
...........
1..........
.11........
.....1.....
...12......
.1....1....`

	lobby.MoveRobots(100)

	testutils.ExpectToMatchString(t, fmt.Sprint(lobby), after100)

	testutils.ExpectToMatchInt(t, len(lobby.RobotsInQuadrants), 4)

	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[0], 1)
	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[1], 3)
	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[2], 4)
	testutils.ExpectToMatchInt(t, lobby.RobotsInQuadrants[3], 1)
}
