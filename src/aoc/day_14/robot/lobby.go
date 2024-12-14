package robot

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"regexp"
)

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("[%d, %d]", c.X, c.Y)
}

type RobotsAtPosition = []*Robot

type Lobby struct {
	XSize             int
	YSize             int
	RobotPositions    map[string]RobotsAtPosition
	robots            []*Robot
	RobotsInQuadrants map[int]int
}

func MakeLobby(XSize int, YSize int) *Lobby {
	return &Lobby{
		XSize:             XSize,
		YSize:             YSize,
		RobotPositions:    make(map[string][]*Robot),
		RobotsInQuadrants: make(map[int]int),
	}
}

func (l Lobby) String() string {
	printer := "\n"
	for y := 0; y < l.YSize; y++ {
		for x := 0; x < l.XSize; x++ {
			coordinate := Coordinate{
				X: x,
				Y: y,
			}
			robotsAt := l.RobotPositions[fmt.Sprint(coordinate)]
			if len(robotsAt) > 0 {
				printer += fmt.Sprint(len(robotsAt))
			} else {
				printer += "."
			}
		}
		if y != l.YSize-1 {
			printer += "\n"
		}
	}
	return printer
}

func (l *Lobby) MoveRobots(time int) {
	// reset robot positions
	l.RobotPositions = make(map[string][]*Robot)
	l.RobotsInQuadrants = make(map[int]int)
	newRobots := make([]*Robot, len(l.robots))
	for i, robot := range l.robots {
		newRobot := robot.CalculatePositionIn(time, l.XSize, l.YSize)
		newRobots[i] = newRobot
		l.PlaceRobot(newRobot)
	}
	l.robots = newRobots
}

func (l *Lobby) PlaceRobot(robot *Robot) {
	coordinate := Coordinate{
		X: robot.X,
		Y: robot.Y,
	}
	robotsAtPoint := l.RobotPositions[fmt.Sprint(coordinate)]
	if robotsAtPoint == nil {
		robotsAtPoint = []*Robot{}
	}
	l.RobotPositions[fmt.Sprint(coordinate)] = append(robotsAtPoint, robot)
	l.assignQuadrant(robot)
}

func (l *Lobby) assignQuadrant(robot *Robot) {
	quadXSize := (l.XSize / 2)
	quadYSize := (l.YSize / 2)

	// ignore robots in center horizontally or vertically
	gapX := l.XSize % (quadXSize)
	gapY := l.YSize % (quadYSize)

	quadrantId := 0
	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			quadMinX := (x * quadXSize) + (x * gapX)
			quadMinY := (y * quadYSize) + (y * gapY)
			quadXMax := quadMinX + quadXSize - 1
			quadYMax := quadMinY + quadYSize - 1

			if robot.X >= quadMinX && robot.X <= quadXMax &&
				robot.Y >= quadMinY && robot.Y <= quadYMax {
				// is in this quadrant
				l.RobotsInQuadrants[quadrantId] += 1
			} else {
				l.RobotsInQuadrants[quadrantId] += 0
			}
			quadrantId++
		}
	}
}

func (l *Lobby) CalculateSafetyFactor() int {
	sum := 0
	for i := 0; i < len(l.RobotsInQuadrants); i++ {
		if l.RobotsInQuadrants[0] == 0 {
			continue
		}
		if sum == 0 {
			sum = l.RobotsInQuadrants[i]
		} else {
			sum *= l.RobotsInQuadrants[i]
		}
	}
	return sum
}

func LoadLobby(path string, XSize int, YSize int) *Lobby {
	file := readfile.Open(path)
	defer file.Close()

	lobby := MakeLobby(XSize, YSize)

	robots := []*Robot{}
	err := readfile.ReadLine(file, func(line string) error {
		re := regexp.MustCompile(`[+|-]*\d+`)
		matches := re.FindAllStringSubmatch(line, -1)
		x := 0
		y := 0
		vX := 0
		vY := 0
		for i, match := range matches {
			number := utils.ToInt(match[0])
			if i == 0 {
				x = number
			} else if i == 1 {
				y = number
			} else if i == 2 {
				vX = number
			} else {
				vY = number
			}
		}
		robot := Robot{
			X: x,
			Y: y,
			Velocity: Velocity{
				X: vX,
				Y: vY,
			},
		}
		lobby.PlaceRobot(&robot)
		robots = append(robots, &robot)
		return nil
	})
	lobby.robots = robots
	if err != nil {
		log.Fatal(err)
	}
	return lobby

}
