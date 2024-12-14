package robot

type Velocity struct {
	X int
	Y int
}

type Robot struct {
	X        int
	Y        int
	Velocity Velocity
}

func (r *Robot) CalculatePositionIn(seconds int, maxX int, maxY int) *Robot {

	xShift := (r.X + (seconds * r.Velocity.X)) % maxX
	yShift := (r.Y + (seconds * r.Velocity.Y)) % maxY

	newX := xShift
	newY := yShift

	if newX < 0 {
		newX = maxX + newX
	}
	if newY < 0 {
		newY = maxY + newY
	}

	return &Robot{
		X:        newX,
		Y:        newY,
		Velocity: r.Velocity,
	}
}
