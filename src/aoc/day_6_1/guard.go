package main

import (
	"errors"
	"fmt"
	"log"
)

type Guard struct {
	Position           *Coordinate
	Direction          GuardDirection
	VisitedCoordinates map[*Coordinate]bool
}

func NewGuard(position *Coordinate, direction GuardDirection) *Guard {
	position.Visited = true
	return &Guard{
		Position:           position,
		Direction:          direction,
		VisitedCoordinates: make(map[*Coordinate]bool),
	}
}

func (g *Guard) TurnRight() {
	switch g.Direction {
	case NORTH:
		g.Direction = EAST
	case EAST:
		g.Direction = SOUTH
	case SOUTH:
		g.Direction = WEST
	case WEST:
		g.Direction = NORTH
	default:
		log.Fatalf("invalid direction")
	}
}

func (g *Guard) CompletePatrol(patrolMap *Map) (bool, error) {
	newCoordinate, err := g.MoveToNextObstacle(patrolMap)
	if newCoordinate == nil {
		log.Println(err)
		// map is completed
		return true, nil
	}
	g.TurnRight()
	return g.CompletePatrol(patrolMap)

}

func (g *Guard) MoveToNextObstacle(patrolMap *Map) (*Coordinate, error) {
	var nextCoordinate *Coordinate
	currX := g.Position.X
	currY := g.Position.Y
	YMax := len(patrolMap.coordinates)
	newX := currX
	newY := currY
	switch g.Direction {
	case NORTH:
		newY = currY - 1
		if newY < 0 {
			g.leave()
			return nil, errors.New("guard left the map")
		}
	case SOUTH:
		newY = currY + 1

		if newY >= YMax {
			g.leave()
			return nil, errors.New("guard left the map")
		}
	case EAST:
		newX = currX + 1
		xMax := len(patrolMap.coordinates[currY])
		if newX >= xMax {
			g.leave()
			return nil, errors.New("guard left the map")
		}
	case WEST:
		newX = currX - 1
		if newX < 0 {
			g.leave()
			return nil, errors.New("guard left the map")
		}
	default:
		log.Fatalf("guard not facing a valid direction")
	}
	nextCoordinate = patrolMap.coordinates[newY][newX]
	if nextCoordinate.Obstacle {
		return g.Position, nil
	}
	g.visit(nextCoordinate)
	return g.MoveToNextObstacle(patrolMap)
}

func (g Guard) String() string {
	return fmt.Sprint(g.Direction)
}

func (g *Guard) visit(c *Coordinate) {
	currentPosition := g.Position
	if currentPosition != nil {
		currentPosition.Guard = nil
	}
	c.Guard = g
	g.Position = c
	c.Visited = true
	g.VisitedCoordinates[c] = true

}

func (g *Guard) leave() {
	currentPosition := g.Position
	currentPosition.Guard = nil
	g.Position = nil

}

type GuardDirection int

// Define directions using iota
const (
	NORTH GuardDirection = iota
	EAST
	SOUTH
	WEST
)

// String returns the string representation of the Direction
func (d GuardDirection) String() string {
	switch d {
	case NORTH:
		return "^"
	case EAST:
		return ">"
	case SOUTH:
		return "v"
	case WEST:
		return "<"
	default:
		return "[INVALID]"
	}
}

func GetGuardDirection(input string) (GuardDirection, error) {
	switch input {
	case "^":
		return NORTH, nil
	case ">":
		return EAST, nil
	case "v":
		return SOUTH, nil
	case "<":
		return WEST, nil
	default:
		return -1, fmt.Errorf("input is not a guard direction: %v", input)
	}
}
