package main

type GridCoordinate struct {
	x, y int
}

func (c *GridCoordinate) GetMASTotal(grid [][]string) int {
	currentLetter := c.getLetter(grid)

	if currentLetter != "A" {
		return 0
	}

	columnLength := len(grid)
	rowLength := len(grid[c.y])

	// check out of bounds
	if c.x-1 < 0 || c.y-1 < 0 || c.y+1 > columnLength-1 || c.x+1 > rowLength-1 {

		return 0
	}

	topLeft := c.makeOffset(-1, -1)
	topRight := c.makeOffset(1, -1)
	bottomLeft := c.makeOffset(-1, 1)
	bottomRight := c.makeOffset(1, 1)

	p := Permutation{
		TopLeft:     topLeft.getLetter(grid),
		TopRight:    topRight.getLetter(grid),
		BottomLeft:  bottomLeft.getLetter(grid),
		BottomRight: bottomRight.getLetter(grid),
	}
	return p.check()
}

// Private
func (c *GridCoordinate) makeOffset(xOffset int, yOffset int) *GridCoordinate {
	newCoord := GridCoordinate{
		x: c.x + xOffset,
		y: c.y + yOffset,
	}

	return &newCoord
}

func (c *GridCoordinate) getLetter(grid [][]string) string {
	letter := grid[c.y][c.x]
	return letter
}
