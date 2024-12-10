package topmap

import (
	"aoc/src/aoc/readfile"
	"aoc/src/aoc/utils"
	"fmt"
	"log"
	"strings"
)

type Location struct {
	X      int
	Y      int
	Height int
}

func (current *Location) getSlope(next *Location) int {
	if next == nil || current == nil {
		return -1
	}
	return next.Height - current.Height
}

type TopMap struct {
	TrailHeads []*TrailHead
	Locations  [][]*Location
}

func (t *TopMap) GetTrailHeadSum() int {
	sum := 0
	for _, trailHead := range t.TrailHeads {
		uniqueEndNodes := make(map[*Location]bool)
		paths := t.GetTrailHeadPaths(trailHead)
		for _, path := range paths {
			uniqueEndNodes[path[len(path)-1]] = true
		}
		sum += len(uniqueEndNodes)
	}

	return sum
}

func (t TopMap) String() string {
	printer := ""

	for i, row := range t.Locations {
		for _, location := range row {
			printer += fmt.Sprint(location.Height)
		}
		if i != len(row)-1 {
			printer += "\n"
		}
	}

	return printer
}

func (t *TopMap) getNodeOffset(offsetX int, offSetY int, l *Location) *Location {
	newY := l.Y + offSetY

	if newY >= 0 && newY < len(t.Locations) {
		row := t.Locations[newY]
		newX := l.X + offsetX
		if newX >= 0 && newX < len(row) {
			return row[newX]
		}
	}

	return nil
}

type TrailHead struct {
	Location *Location
}

func PrintPath(topMap *TopMap, trail []*Location) string {
	printer := ""

	for _, location := range trail {
		printer += fmt.Sprintf(" [%d, %d]:%d ", location.X, location.Y, location.Height)
	}

	return printer
}

func ValidateTrail(topMap *TopMap, trail []*Location) bool {
	if len(trail) == 0 {
		return false
	}
	lastNode := trail[len(trail)-1]
	if lastNode.Height != 9 {
		return false
	}
	prevNode := trail[0]
	if prevNode.Height != 0 {
		return false
	}
	for _, location := range trail {
		if location == prevNode {
			// first node
			continue
		}
		if location.Height != prevNode.Height+1 {
			return false
		}
		prevNode = location
	}
	return true
}

func (topMap *TopMap) GetTrailHeadPaths(t *TrailHead) [][]*Location {
	if t.Location.Height != 0 {
		log.Fatal("Cannot calculate paths on non 0 node")
	}

	path := []*Location{}
	return topMap.calculatePaths(t.Location, path)
}

func (topMap *TopMap) calculatePaths(location *Location, path []*Location) [][]*Location {
	if location.Height == 9 {
		path = append(path, location)
		return [][]*Location{path}
	}

	path = append(path, location)

	north := topMap.getNodeOffset(0, -1, location)
	east := topMap.getNodeOffset(1, 0, location)
	west := topMap.getNodeOffset(-1, 0, location)
	south := topMap.getNodeOffset(0, 1, location)

	var allPaths [][]*Location

	if north != nil && location.getSlope(north) == 1 {
		allPaths = append(allPaths, topMap.calculatePaths(north, append([]*Location{}, path...))...)
	}
	if east != nil && location.getSlope(east) == 1 {
		allPaths = append(allPaths, topMap.calculatePaths(east, append([]*Location{}, path...))...)
	}
	if west != nil && location.getSlope(west) == 1 {
		allPaths = append(allPaths, topMap.calculatePaths(west, append([]*Location{}, path...))...)
	}
	if south != nil && location.getSlope(south) == 1 {
		allPaths = append(allPaths, topMap.calculatePaths(south, append([]*Location{}, path...))...)
	}

	return allPaths
}

func LoadMap(path string) *TopMap {
	file := readfile.Open(path)
	defer file.Close()
	rowCounter := 0
	locations := [][]*Location{}
	trailHeads := []*TrailHead{}
	err := readfile.ReadLine(file, func(line string) error {
		numsStr := strings.Split(line, "")
		row := []*Location{}
		// process file and block space at once
		for i := 0; i < len(numsStr); i++ {
			height := utils.ToInt(numsStr[i])
			location := Location{
				Height: height,
				X:      i,
				Y:      rowCounter,
			}
			if location.Height == 0 {
				trailHeads = append(trailHeads, &TrailHead{
					Location: &location,
				})
			}
			row = append(row, &location)
		}
		locations = append(locations, row)
		rowCounter++
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return &TopMap{
		TrailHeads: trailHeads,
		Locations:  locations,
	}
}
