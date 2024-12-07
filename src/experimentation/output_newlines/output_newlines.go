package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Static text at the start with color
	fmt.Println("\033[1;34mWelcome to the Task Progress Tracker!\033[0m") // Blue
	fmt.Println("\033[1;33mHere are the current statuses:\033[0m")        // Yellow

	// Initial dynamic multiline string with colored task labels
	output := []string{
		"\033[1;32mTask 1:\033[0m [          ] 0%",
		"\033[1;32mTask 2:\033[0m [          ] 0%",
		"\033[1;32mTask 3:\033[0m [          ] 0%",
	}

	// Print each line of the initial output
	for _, line := range output {
		fmt.Println(line)
	}

	// Simulate dynamic updates for progress bars
	for i := 0; i <= 10; i++ {
		// Clear previous lines
		clearLines(len(output))

		// Update the output with progress bars and task labels
		output = []string{
			fmt.Sprintf("\033[1;32mTask 1:\033[0m [%s] %d%%", progressBar(4+i), 40+i*6),
			fmt.Sprintf("\033[1;32mTask 2:\033[0m [%s] %d%%", progressBar(2+i), 20+i*8),
			fmt.Sprintf("\033[1;32mTask 3:\033[0m [%s] %d%%", progressBar(i), i*10),
		}

		// Print each updated line
		for _, line := range output {
			fmt.Println(line)
		}

		// Pause for a while to simulate the progress
		time.Sleep(500 * time.Millisecond)
	}

	// Print the completion message with color
	fmt.Println("\033[1;32mAll tasks completed!\033[0m") // Green
}

// Helper to sanitize strings (remove extra spaces or newlines)
func sanitizeString(s string) string {
	return strings.TrimSpace(s)
}

// Helper to clear lines above the cursor (used for updating)
func clearLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print("\033[A\033[2K") // Move up and clear the line
	}
}

// Helper to generate a progress bar string with green hashes
func progressBar(progress int) string {
	total := 10
	bar := ""
	for i := 0; i < total; i++ {
		if i < progress {
			bar += "\033[1;32m#\033[0m" // Green for filled progress
		} else {
			bar += " " // Empty space for unfilled progress
		}
	}
	return bar
}
