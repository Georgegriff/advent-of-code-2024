package maze

import (
	"aoc/src/aoc/dijkstra"
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
	"strings"
)

type Coordinate struct {
	X       int
	Y       int
	isStart bool
	isWall  bool
	isEnd   bool
}

func (c Coordinate) String() string {
	return fmt.Sprintf("[%d, %d]", c.X, c.Y)
}

func NewCoordinate(X, Y int, character string) *Coordinate {
	isWall := false
	isStart := false
	isEnd := false
	if character == "#" {
		isWall = true
	} else if character == "S" {
		isStart = true
	} else if character == "E" {
		isEnd = true
	}

	return &Coordinate{
		X:       X,
		Y:       Y,
		isStart: isStart,
		isEnd:   isEnd,
		isWall:  isWall,
	}
}

type Maze struct {
	coordinates    [][]*Coordinate
	Start          *Coordinate
	End            *Coordinate
	startDirection Direction
}

func (m *Maze) getCoordinateInDirection(coord *Coordinate, direction Direction) *Coordinate {
	coordinateOffset := Coordinate{
		X: 0,
		Y: 0,
	}
	switch direction {
	case WEST:
		coordinateOffset.X = -1
	case EAST:
		coordinateOffset.X = 1
	case SOUTH:
		coordinateOffset.Y = 1
	case NORTH:
		coordinateOffset.Y = -1
	}
	newX := coord.X + coordinateOffset.X
	newY := coord.Y + coordinateOffset.Y
	maxY := len(m.coordinates) - 1
	if newY < 0 || newY > maxY {
		return nil
	}
	maxX := len(m.coordinates[newY]) - 1
	if newX < 0 || newX > maxX {
		return nil
	}
	return m.coordinates[newY][newX]
}

func (m Maze) String() string {
	mapPrinter := "\n"
	for i, row := range m.coordinates {
		for _, coord := range row {
			if coord.isWall {
				mapPrinter += "#"
			} else if coord.isStart {
				mapPrinter += "S"
			} else if coord.isEnd {
				mapPrinter += "E"
			} else {
				mapPrinter += "."
			}
		}
		if i != len(m.coordinates)-1 {
			mapPrinter += "\n"
		}
	}
	return mapPrinter
}

func LoadMaze(path string, startDirection Direction) *Maze {
	file := readfile.Open(path)
	defer file.Close()
	rowCounter := 0
	coordinates := [][]*Coordinate{}
	var start *Coordinate
	var end *Coordinate
	err := readfile.ReadLine(file, func(line string) error {
		positionOnLine := strings.Split(line, "")

		row := []*Coordinate{}
		for x, character := range positionOnLine {
			coord := NewCoordinate(x, rowCounter, character)
			if coord.isStart {
				start = coord
			}
			if coord.isEnd {
				end = coord
			}
			row = append(row, coord)
		}
		coordinates = append(coordinates, row)
		rowCounter++
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &Maze{
		coordinates:    coordinates,
		Start:          start,
		End:            end,
		startDirection: startDirection,
	}
}

func (m *Maze) GetBestPathScore() int {
	fmt.Println("Building graph...")
	graph := m.BuildGraph()
	start := fmt.Sprintf("%v", m.Start)
	end := fmt.Sprintf("%v", m.End)
	fmt.Println("Built graph, finding best path...")
	cost, best := graph.GetPath(start, end)
	fmt.Printf("\n%v", strings.Join(best, " "))
	return cost
}

func GetNextDirection(curr *Coordinate, next *Coordinate) Direction {
	if next.X > curr.X {
		return EAST
	} else if next.X < curr.X {
		return WEST
	} else if next.Y > curr.Y {
		return SOUTH
	} else if next.Y < curr.Y {
		return NORTH
	}

	return -1
}

func (m *Maze) BuildGraph() *dijkstra.Graph {
	graph := dijkstra.MakeGraph()
	visited := make(map[string]bool)
	m.findEdges(m.Start, m.startDirection, visited, graph)
	return graph
}

func (m *Maze) findEdges(current *Coordinate, currentDirection Direction, visited map[string]bool, graph *dijkstra.Graph) {
	if current == nil || current.isWall || current.isEnd {
		return
	}

	currentKey := fmt.Sprintf("%v", current)
	if visited[currentKey] {
		return
	}

	visited[currentKey] = true

	for _, direction := range []Direction{NORTH, EAST, SOUTH, WEST} {
		next := m.getCoordinateInDirection(current, direction)
		if next == nil || next.isWall {
			continue
		}

		nextNode := fmt.Sprint(next)

		if direction != currentDirection {

			graph.AddEdge(currentKey, nextNode, 1001)
		} else {
			graph.AddEdge(currentKey, nextNode, 1)

		}
		m.findEdges(next, direction, visited, graph)
	}

}

type Direction int

// Define directions using iota
const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d Direction) String() string {
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

func (d Direction) GoString() string {
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

func GetDirection(input string) Direction {
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
		log.Fatalf("not a valid  direction %s", input)
	}
	return -1
}
