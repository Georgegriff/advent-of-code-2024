package main

import (
	"fmt"
	"time"
)

func main() {
	// Static text at the start
	fmt.Println("Welcome to the Task Progress Tracker!")
	fmt.Println("Tracking the following tasks:")

	// Initial multiline dynamic content
	lines := []string{
		"Task 1: [          ] 0%",
		"Task 2: [          ] 0%",
		"Task 3: [          ] 0%",
	}

	// Print initial lines
	for _, line := range lines {
		fmt.Println(line)
	}

	// Start updating the dynamic content
	for i := 0; i <= 10; i++ {
		// Move cursor up to the starting line of dynamic content
		moveCursorUp(len(lines))

		// Update progress
		lines[0] = fmt.Sprintf("Task 1: [%s] %d%%", progressBar(4+i), 40+i*6)
		lines[1] = fmt.Sprintf("Task 2: [%s] %d%%", progressBar(2+i), 20+i*8)
		lines[2] = fmt.Sprintf("Task 3: [%s] %d%%", progressBar(i), i*10)

		// Reprint updated lines
		for _, line := range lines {
			fmt.Println(line)
		}

		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("All tasks completed!")
}

// Helper to move cursor up a specific number of lines
func moveCursorUp(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[A") // Move up one line
	}
}

// Helper to generate a progress bar string
func progressBar(progress int) string {
	total := 10
	bar := ""
	for i := 0; i < total; i++ {
		if i < progress {
			bar += "#"
		} else {
			bar += " "
		}
	}
	return bar
}
