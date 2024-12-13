package garden

import (
	"aoc/src/aoc/colors"
	"aoc/src/aoc/readfile"
	"fmt"
	"log"
	"slices"
	"strings"
)

type RegionType = string

type Garden struct {
	regions     []*Region
	regionTypes []RegionType
	plots       [][]*Plot
}

var plotColors []colors.Color = colors.AllColors()

func (g *Garden) VisitConnectedPlots(p *Plot, xMax int, yMax int, region *Region) {

	if p.region != nil {
		// already visited
		return
	}
	if region == nil {
		plots := []*Plot{p}
		p.region = &Region{
			RegionType: p.RegionType,
			plots:      plots,
		}
		g.regions = append(g.regions, p.region)
	} else {
		if region.RegionType != p.RegionType {
			return
		}
		p.region = region
		region.plots = append(region.plots, p)
	}
	// check north
	if p.Y-1 >= 0 {
		newPlot := g.plots[p.Y-1][p.X]
		g.VisitConnectedPlots(newPlot, xMax, yMax, p.region)
	}
	// check east
	if p.X+1 <= xMax {
		newPlot := g.plots[p.Y][p.X+1]
		g.VisitConnectedPlots(newPlot, xMax, yMax, p.region)
	}
	// check west
	if p.X-1 >= 0 {
		newPlot := g.plots[p.Y][p.X-1]
		g.VisitConnectedPlots(newPlot, xMax, yMax, p.region)
	}
	// check south
	if p.Y+1 <= yMax {
		newPlot := g.plots[p.Y+1][p.X]
		g.VisitConnectedPlots(newPlot, xMax, yMax, p.region)
	}
}

func (g Garden) String() string {
	printer := "\n"
	for i, row := range g.plots {
		for _, coord := range row {
			rType := coord.RegionType

			colorIndex := IndexOf(g.regionTypes, rType)
			color := plotColors[colorIndex]
			printer += colors.PrintColor(coord.RegionType, color)
		}
		if i != len(g.plots)-1 {
			printer += "\n"
		}
	}

	return printer
}

func (g *Garden) visitGarden() {
	for _, row := range g.plots {
		for _, plot := range row {
			g.VisitConnectedPlots(plot, len(g.plots)-1, len(row)-1, plot.region)
		}
	}
}

func (g *Garden) GetPrice() int {
	sum := 0
	for _, region := range g.regions {
		sum += region.GetArea() * region.GetPerimeter(g)
	}
	return sum
}

func LoadGarden(path string) *Garden {

	file := readfile.Open(path)
	defer file.Close()
	rowCounter := 0
	plots := [][]*Plot{}
	regionTypes := []RegionType{}
	regions := []*Region{}
	err := readfile.ReadLine(file, func(line string) error {
		positionOnLine := strings.Split(line, "")

		row := []*Plot{}
		for x, regionType := range positionOnLine {
			if !slices.Contains(regionTypes, regionType) {
				regionTypes = append(regionTypes, regionType)
			}
			coord := Plot{
				RegionType: regionType,
				X:          x,
				Y:          rowCounter,
			}

			row = append(row, &coord)
		}
		plots = append(plots, row)
		rowCounter++
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	g := Garden{
		plots:       plots,
		regionTypes: regionTypes,
		regions:     regions,
	}
	g.visitGarden()
	return &g
}

type Region struct {
	RegionType RegionType
	plots      []*Plot
}

func (r *Region) GetArea() int {
	return len(r.plots)
}

func (r *Region) isBorder(g *Garden, p *Plot, xOffset, yOffset, xMax, yMax int) bool {
	if p.Y+yOffset >= 0 && p.X+xOffset >= 0 && p.Y+yOffset <= yMax && p.X+xOffset <= xMax {
		newPlot := g.plots[p.Y+yOffset][p.X+xOffset]
		if !slices.Contains(r.plots, newPlot) {
			return true
		}
	} else {
		return true
	}
	return false
}

func (r *Region) isInRegion(g *Garden, p *Plot, xOffset, yOffset, xMax, yMax int) bool {
	if p.Y+yOffset >= 0 && p.X+xOffset >= 0 && p.Y+yOffset <= yMax && p.X+xOffset <= xMax {
		newPlot := g.plots[p.Y+yOffset][p.X+xOffset]
		if slices.Contains(r.plots, newPlot) {
			return true
		}
	} else {
		return false
	}
	return false
}

func (r *Region) GetPerimeter(g *Garden) int {
	fencePlots := []*Plot{}
	yMax := len(g.plots) - 1
	for _, p := range r.plots {
		xMax := len(g.plots[0]) - 1

		topIsBorder := r.isBorder(g, p, 0, -1, xMax, yMax)
		leftIsBorder := r.isBorder(g, p, -1, 0, xMax, yMax)
		rightIsBorder := r.isBorder(g, p, 1, 0, xMax, yMax)
		bottomIsBorder := r.isBorder(g, p, 0, 1, xMax, yMax)
		diagRightDown := r.isInRegion(g, p, 1, 1, xMax, yMax)
		diagLeftDown := r.isInRegion(g, p, -1, 1, xMax, yMax)
		diagRightUp := r.isInRegion(g, p, 1, -1, xMax, yMax)
		diagLeftUp := r.isInRegion(g, p, -1, -1, xMax, yMax)

		// corners
		if topIsBorder && leftIsBorder && !diagLeftUp {
			fencePlots = append(fencePlots, p)
		}
		if topIsBorder && rightIsBorder && !diagRightUp {
			fencePlots = append(fencePlots, p)
		}
		if leftIsBorder && bottomIsBorder && !diagLeftDown {
			fencePlots = append(fencePlots, p)
		}
		if rightIsBorder && bottomIsBorder && !diagRightDown {
			fencePlots = append(fencePlots, p)
		}
		// diag down right
		if rightIsBorder && diagRightDown && !bottomIsBorder {
			fencePlots = append(fencePlots, p)
		}

		// diag down left
		if leftIsBorder && diagLeftDown && !bottomIsBorder {
			fencePlots = append(fencePlots, p)
		}
		if bottomIsBorder && diagRightDown && !rightIsBorder {
			fencePlots = append(fencePlots, p)
		}
		if bottomIsBorder && diagLeftDown && !leftIsBorder {
			fencePlots = append(fencePlots, p)
		}

		// obtuse corners
		if topIsBorder && rightIsBorder && diagRightUp {
			fencePlots = append(fencePlots, p)
		}
		if topIsBorder && leftIsBorder && diagLeftUp {
			fencePlots = append(fencePlots, p)
		}
		if bottomIsBorder && rightIsBorder && diagRightDown {
			fencePlots = append(fencePlots, p)
		}
		if bottomIsBorder && leftIsBorder && diagLeftDown {
			fencePlots = append(fencePlots, p)
		}
	}

	return len(fencePlots)
}

func (r Region) String() string {
	return fmt.Sprintf("\n%s: plots %v\n", r.RegionType, r.plots)
}

type Plot struct {
	X          int
	Y          int
	RegionType RegionType
	region     *Region
}

func (r Plot) String() string {
	return fmt.Sprintf("[%d, %d]", r.X, r.Y)
}

func IndexOf(slice []string, element string) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}
