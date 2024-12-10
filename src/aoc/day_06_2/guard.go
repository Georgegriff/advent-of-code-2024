package main

import (
	"fmt"
	"log"
)

type VisitedCoordinates map[*Coordinate]map[GuardDirection]bool

type Guard struct {
	Position           *Coordinate
	Direction          GuardDirection
	VisitedCoordinates VisitedCoordinates
}

func NewGuard(position *Coordinate, direction GuardDirection) *Guard {
	position.Visited = true
	if direction == NORTH || direction == SOUTH {
		position.VisitedType = VERTICAL
	} else {
		position.VisitedType = HORIZONTAL
	}
	return &Guard{
		Position:           position,
		Direction:          direction,
		VisitedCoordinates: make(VisitedCoordinates),
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
		if err != nil {
			switch e := err.(type) {
			case *CycleError:
				patrolMap.PrintMapState()
				return false, e
			}
		}
		// log.Println(err)
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
			err := &LeftMapError{Coord: g.Position}
			g.Leave()
			return nil, err
		}
	case SOUTH:
		newY = currY + 1

		if newY >= YMax {
			err := &LeftMapError{Coord: g.Position}
			g.Leave()
			return nil, err
		}
	case EAST:
		newX = currX + 1
		xMax := len(patrolMap.coordinates[currY])
		if newX >= xMax {
			err := &LeftMapError{Coord: g.Position}
			g.Leave()
			return nil, err
		}
	case WEST:
		newX = currX - 1
		if newX < 0 {
			err := &LeftMapError{Coord: g.Position}
			g.Leave()
			return nil, err
		}
	default:
		log.Fatalf("guard not facing a valid direction")
	}
	nextCoordinate = patrolMap.coordinates[newY][newX]
	if nextCoordinate.Obstacle {
		return g.Position, nil
	}
	coordinateVisits := g.VisitedCoordinates[nextCoordinate]
	if coordinateVisits != nil && coordinateVisits[g.Direction] {
		return nil, &CycleError{Coord: nextCoordinate}
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
	if g.VisitedCoordinates[c] == nil {
		g.VisitedCoordinates[c] = make(map[GuardDirection]bool)

	}
	g.VisitedCoordinates[c][g.Direction] = true

}

func (g *Guard) Leave() {
	currentPosition := g.Position
	if currentPosition != nil {
		currentPosition.Guard = nil
		g.Position = nil
	}
}

func (g *Guard) Join(position *Coordinate, direction GuardDirection) {
	position.Guard = g
	g.Position = position
	g.Direction = direction
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
		return "\u2191" // ↑
	case EAST:
		return "\u2192" // ↓
	case SOUTH:
		return "\u2193" // →
	case WEST:
		return "\u2190" // ←
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

type LeftMapError struct {
	Coord *Coordinate
}

func (e *LeftMapError) Error() string {
	return fmt.Sprintf("coordinate %v is outside the map to the left", e.Coord)
}

// CycleError represents an error related to cycling within coordinates
type CycleError struct {
	Coord *Coordinate
}

func (e *CycleError) Error() string {
	return fmt.Sprintf("cycle detected at coordinate %v", e.Coord)
}
