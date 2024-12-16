package warehouse

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
	"slices"
	"strings"
)

const ROBOT = "@"
const BOX = "O"
const WALL = "#"
const UP_DIRECTION = "^"
const DOWN_DIRECTION = "v"
const RIGHT_DIRECTION = ">"
const LEFT_DIRECTION = "<"

/** Generics */
type Entity interface {
	GetType() string
	String() string
}

/** Position */
type Position struct {
	X      int
	Y      int
	entity Entity
}

func (p *Position) EntityType() string {
	if p.entity == nil {
		return ""
	}
	return p.entity.GetType()
}

/** Box */
type Box struct {
	start *Position
	end   *Position
}

func (b *Box) GetType() string {
	return BOX
}

func (b Box) String() string {
	return b.GetType()
}

func MakeBox(start *Position, end *Position) *Box {
	box := Box{
		start: start,
		end:   end,
	}
	start.entity = &box
	end.entity = &box

	return &box
}

func (b *Box) updatePosition(newStart *Position, newEnd *Position) {
	b.start.entity = nil
	b.end.entity = nil
	newStart.entity = b
	newEnd.entity = b
	b.start = newStart
	b.end = newEnd
}

func (b *Box) GetGPS(maxX, maxY int) int {
	// maxXDiff := maxX - b.end.X
	x := b.start.X
	y := b.end.Y
	// if maxXDiff < b.start.X {
	// 	x = b.end.X
	// }

	return (100 * y) + x
}

/** Wall */
type Wall struct {
	start *Position
	end   *Position
}

func (w *Wall) GetType() string {
	return WALL
}

func (w Wall) String() string {
	return w.GetType()
}

func MakeWall(start *Position, end *Position) *Wall {
	wall := Wall{
		start: start,
		end:   end,
	}
	start.entity = &wall
	end.entity = &wall

	return &wall
}

/** Robot */
type Robot struct {
	Instructions       string
	CurrentInstruction int
	*Position
}

func (b *Robot) GetType() string {
	return ROBOT
}

func (r Robot) String() string {
	return r.GetType()
}

func (r *Robot) GetDirection() RobotDirection {
	instructions := []rune(r.Instructions)
	instruction := GetRobotDirection(string(instructions[r.CurrentInstruction]))
	return instruction
}

func (r *Robot) Move(positions [][]*Position) bool {
	r.moveInDirection(r.GetDirection(), positions)
	if r.CurrentInstruction+1 <= len(r.Instructions)-1 {
		r.CurrentInstruction++
		return false
	}
	// finished instructions
	return true

}

func MakeNewPosition(direction RobotDirection, current *Position, positions [][]*Position) *Position {
	positionOffset := Position{
		X: 0,
		Y: 0,
	}
	switch direction {
	case NORTH:
		positionOffset.Y = -1
	case EAST:
		positionOffset.X = 1
	case WEST:
		positionOffset.X = -1
	case SOUTH:
		positionOffset.Y = 1
	}
	newX := current.X + positionOffset.X
	newY := current.Y + positionOffset.Y
	if newY < 0 || newY > len(positions)-1 {
		return nil
	}
	row := positions[newY]
	if newX < 0 || newX > len(row)-1 {
		return nil
	}
	return row[newX]
}

func SortBoxesByDirection(boxes []*Box, direction RobotDirection) {
	switch direction {
	case NORTH:
		slices.SortFunc(boxes, func(a, b *Box) int {
			if a.start.Y < b.start.Y {
				return -1
			}
			if a.start.Y > b.start.Y {
				return 1
			}
			return 0
		})
	case EAST:
		slices.SortFunc(boxes, func(a, b *Box) int {
			if a.start.X > b.start.X {
				return -1
			}
			if a.start.X < b.start.X {
				return 1
			}
			return 0
		})
	case SOUTH:
		slices.SortFunc(boxes, func(a, b *Box) int {
			if a.start.Y > b.start.Y {
				return -1
			}
			if a.start.Y < b.start.Y {
				return 1
			}
			return 0
		})
	case WEST:
		slices.SortFunc(boxes, func(a, b *Box) int {
			if a.start.X < b.start.X {
				return -1
			}
			if a.start.X > b.start.X {
				return 1
			}
			return 0
		})
	}
}

func (r *Robot) moveInDirection(direction RobotDirection, positions [][]*Position) {
	newPosition := MakeNewPosition(direction, r.Position, positions)
	if newPosition == nil {
		return
	}
	if newPosition.entity == nil {
		// can safely move the robot
		r.updatePosition(newPosition)
		return
	}
	if newPosition.entity.GetType() == WALL {
		return
	}
	if newPosition.entity.GetType() == BOX {
		// figure out if all the connected boxes can be moved
		boxesToMove, canMoveBoxes := r.canMoveBoxes(direction, newPosition, positions)
		if canMoveBoxes && len(boxesToMove) > 0 {
			// Sort boxes based on direction
			SortBoxesByDirection(boxesToMove, direction)
			// Move boxes in sorted order
			for _, box := range boxesToMove {
				startPositionNext := MakeNewPosition(direction, box.start, positions)
				endPositionNext := MakeNewPosition(direction, box.end, positions)
				box.updatePosition(startPositionNext, endPositionNext)
			}
			// we freed up space for the robot "robot pushed the boxes"
			r.updatePosition(newPosition)
		}
		return
	}
}

func (r *Robot) updatePosition(newPosition *Position) {
	r.Position.entity = nil
	r.Position = newPosition
	newPosition.entity = r
}

func (r *Robot) canMoveBoxes(direction RobotDirection, newPosition *Position, positions [][]*Position) ([]*Box, bool) {
	if newPosition == nil {
		return nil, false
	}
	if newPosition.entity == nil {
		return []*Box{}, true
	}
	box := newPosition.entity.(*Box)
	var boxes = []*Box{box}
	var nextPosition *Position
	if direction == EAST || direction == WEST {
		if direction == WEST {
			nextPosition = MakeNewPosition(direction, box.start, positions)

		} else if direction == EAST {
			nextPosition = MakeNewPosition(direction, box.end, positions)
		}
		if nextPosition.entity == nil {
			return boxes, true
		}
		if nextPosition.entity.GetType() == WALL {
			return nil, false
		}

	} else {
		startPosition := MakeNewPosition(direction, box.start, positions)
		endPosition := MakeNewPosition(direction, box.end, positions)
		if startPosition.entity == nil && endPosition.entity == nil && startPosition.entity == endPosition.entity {
			// box := startPosition.entity.(*Box)
			// boxes = append(boxes, box)
			return boxes, true
		}
		if (startPosition.entity != nil && startPosition.entity.GetType() == WALL) || (endPosition.entity != nil && endPosition.entity.GetType() == WALL) {
			return nil, false
		}
	}

	// we have hit a box
	if direction == EAST || direction == WEST {
		childBoxes, canAdd := r.canMoveBoxes(direction, nextPosition, positions)
		if !canAdd {
			return nil, false
		}
		for _, childBox := range childBoxes {
			if !slices.Contains(boxes, childBox) {
				boxes = append(boxes, childBox)
			}
		}
		return boxes, true
	} else {
		startPosition := MakeNewPosition(direction, box.start, positions)
		endPosition := MakeNewPosition(direction, box.end, positions)
		var startBox *Box
		var endBox *Box
		if startPosition.entity != nil {
			startBox = startPosition.entity.(*Box)
		}
		if endPosition.entity != nil {

			endBox = endPosition.entity.(*Box)
		}

		if startBox == endBox {
			if !slices.Contains(boxes, startBox) {
				boxes = append(boxes, startBox)
			}
		} else {
			if startBox != nil && !slices.Contains(boxes, startBox) {
				boxes = append(boxes, startBox)
			}
			if endBox != nil && !slices.Contains(boxes, endBox) {
				boxes = append(boxes, endBox)
			}
		}
		childBoxesStart, canAddStart := r.canMoveBoxes(direction, startPosition, positions)
		if !canAddStart {
			return nil, false
		}

		childBoxesEnd, canAddEnd := r.canMoveBoxes(direction, endPosition, positions)
		if !canAddEnd {
			return nil, false
		}
		for _, childStartBox := range childBoxesStart {
			if !slices.Contains(boxes, childStartBox) {
				boxes = append(boxes, childStartBox)
			}
		}
		for _, childEndBox := range childBoxesEnd {
			if !slices.Contains(boxes, childEndBox) {
				boxes = append(boxes, childEndBox)

			}

		}

		return boxes, true

	}

}

func MakeRobot(instructions string, initialPosition *Position) *Robot {
	robot := &Robot{
		Instructions:       instructions,
		CurrentInstruction: 0,
		Position:           initialPosition,
	}
	initialPosition.entity = robot

	return robot
}

type RobotDirection int

// Define directions using iota
const (
	NORTH RobotDirection = iota
	EAST
	SOUTH
	WEST
)

func (d RobotDirection) String() string {
	switch d {
	case NORTH:
		return "^"
	case EAST:
		return ">"
	case SOUTH:
		return "V"
	case WEST:
		return "<"
	default:
		return "[INVALID]"
	}
}

func (d RobotDirection) GoString() string {
	switch d {
	case NORTH:
		return "^"
	case EAST:
		return ">"
	case SOUTH:
		return "V"
	case WEST:
		return "<"
	default:
		return "[INVALID]"
	}
}

func GetRobotDirection(input string) RobotDirection {
	switch input {
	case "^":
		return NORTH
	case ">":
		return EAST
	case "v":
		return SOUTH
	case "<":
		return WEST
	default:
		log.Fatalf("not a valid robot direction %s", input)
	}
	return -1
}

/** Warehouse */
type Warehouse struct {
	positions [][]*Position
	boxes     []*Box
	robot     *Robot
}

func (w Warehouse) String() string {
	warehousePrinter := "\n"
	for i, row := range w.positions {
		for _, position := range row {
			if position.entity != nil {
				if box, ok := position.entity.(*Box); ok {
					if box.start == position {
						warehousePrinter += "["
					} else {
						warehousePrinter += "]"
					}
				} else {
					warehousePrinter += fmt.Sprint(position.entity)
				}

			} else {
				warehousePrinter += "."
			}
		}
		if i != len(w.positions)-1 {
			warehousePrinter += "\n"
		}
	}
	return warehousePrinter
}

func (w *Warehouse) Next() bool {

	isFinished := w.robot.Move(w.positions)
	return isFinished
}

func (w *Warehouse) Process() {
	for {
		if w.Next() {
			// fmt.Printf("\n%v", w.robot.GetDirection())
			break
		}
	}
}

func (w *Warehouse) GetGPSSum() int {
	sum := 0

	for _, box := range w.boxes {
		maxY := len(w.positions)
		maxX := len(w.positions[box.start.Y])
		sum += box.GetGPS(maxX, maxY)
	}

	return sum
}

func LoadWarehouse(path string, scale int) *Warehouse {
	file := readfile.Open(path)
	defer file.Close()
	rowCounter := 0
	positions := [][]*Position{}
	var robotPosition *Position
	boxes := []*Box{}
	robotInstructions := ""
	processingState := "warehouse"
	err := readfile.ReadLine(file, func(line string) error {
		if line == "" {
			processingState = "instructions"
		}
		if processingState == "warehouse" {
			positionsOnLine := strings.Split(line, "")
			row := []*Position{}
			columnCounter := 0
			for _, entityType := range positionsOnLine {
				rowPos := []*Position{}
				for i := 0; i < scale; i++ {
					pos := Position{
						X: columnCounter + i,
						Y: rowCounter,
					}
					rowPos = append(rowPos, &pos)
				}
				columnCounter += scale

				if entityType == WALL {
					MakeWall(rowPos[0], rowPos[len(rowPos)-1])
				} else if entityType == BOX {
					box := MakeBox(rowPos[0], rowPos[len(rowPos)-1])
					boxes = append(boxes, box)
				} else if entityType == ROBOT {
					robotPosition = rowPos[0]
				}
				row = append(row, rowPos...)
			}
			positions = append(positions, row)
			rowCounter++
		} else if processingState == "instructions" {
			robotInstructions += line
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	var robot *Robot
	if robotPosition != nil {
		robot = MakeRobot(robotInstructions, robotPosition)
	}

	return &Warehouse{
		positions: positions,
		robot:     robot,
		boxes:     boxes,
	}
}
