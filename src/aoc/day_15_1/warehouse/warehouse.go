package warehouse

import (
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
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

type Moveable interface {
	updatePosition(p *Position)
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
	*Position
	Moveable
}

func (b *Box) GetType() string {
	return BOX
}

func (b Box) String() string {
	return b.GetType()
}

func MakeBox(position *Position) *Box {
	box := Box{
		Position: position,
	}
	position.entity = &box

	return &box
}

func (b *Box) updatePosition(newPosition *Position) {
	b.Position.entity = nil
	b.Position = newPosition
	newPosition.entity = b
}

func (b *Box) GetGPS() int {
	return (100 * b.Y) + b.X
}

/** Wall */
type Wall struct {
	*Position
}

func (w *Wall) GetType() string {
	return WALL
}

func (w Wall) String() string {
	return w.GetType()
}

func MakeWall(position *Position) *Wall {
	wall := Wall{
		Position: position,
	}
	position.entity = &wall

	return &wall
}

/** Robot */
type Robot struct {
	Instructions       string
	CurrentInstruction int
	*Position
	Moveable
}

func (b *Robot) GetType() string {
	return ROBOT
}

func (r Robot) String() string {
	return r.GetType()
}

func (r *Robot) Move(positions [][]*Position) bool {
	instructions := []rune(r.Instructions)
	instruction := GetRobotDirection(string(instructions[r.CurrentInstruction]))
	r.moveInDirection(instruction, positions)
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
		r.moveBoxes(direction, newPosition, positions)
		// we freed up space for the robot "robot pushed the boxes"
		if newPosition.entity == nil {
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

func (r *Robot) moveBoxes(direction RobotDirection, firstBoxPosition *Position, positions [][]*Position) {

	if box, ok := firstBoxPosition.entity.(*Box); ok {
		if !ok {
			return
		}
		// if the next position is empty move there
		newPosition := MakeNewPosition(direction, box.Position, positions)
		if newPosition == nil {
			return
		}
		if newPosition.entity == nil {
			box.updatePosition(newPosition)
			return
		} else if newPosition.entity.GetType() == WALL {
			return
		} else if newPosition.entity.GetType() == BOX {
			r.moveBoxes(direction, newPosition, positions)
			if newPosition.entity == nil {
				box.updatePosition(newPosition)
			}
		}
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
				warehousePrinter += fmt.Sprint(position.entity)
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
	return w.robot.Move(w.positions)
}

func (w *Warehouse) Process() {
	for {
		if w.Next() {
			break
		}
	}
}

func (w *Warehouse) GetGPSSum() int {
	sum := 0

	for _, box := range w.boxes {
		sum += box.GetGPS()
	}

	return sum
}

func LoadWarehouse(path string) *Warehouse {
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
			for x, entityType := range positionsOnLine {
				pos := Position{
					X: x,
					Y: rowCounter,
				}

				if entityType == WALL {
					MakeWall(&pos)
				} else if entityType == BOX {
					box := MakeBox(&pos)
					boxes = append(boxes, box)
				} else if entityType == ROBOT {
					robotPosition = &pos
				}
				row = append(row, &pos)
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

	robot := MakeRobot(robotInstructions, robotPosition)

	return &Warehouse{
		positions: positions,
		robot:     robot,
		boxes:     boxes,
	}
}
