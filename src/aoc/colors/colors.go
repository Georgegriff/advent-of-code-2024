package colors

import "fmt"

var purple = "\033[35m" // Purple
var green = "\033[32m"  // Green
var red = "\033[31m"    // Red
var yellow = "\033[33m" // Yellow
var reset = "\033[0m"   // Reset
var gray = "\033[90m"   // Gray

type OnReadLine func(line string) error

func PrintColor(input string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, input, reset)
}

type Color int

// Define directions using iota
const (
	PURPLE Color = iota
	GREEN
	RED
	YELLOW
	GRAY
)

// String returns the string representation of the Direction
func (d Color) String() string {
	switch d {
	case PURPLE:
		return purple
	case GREEN:
		return green
	case RED:
		return red
	case YELLOW:
		return yellow
	case GRAY:
		return gray
	default:
		return "[[INVALID COLOR]]"
	}
}