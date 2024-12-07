package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 100; i++ {
		// Print the progress percentage with a carriage return (\r)
		fmt.Printf("\rProgress: %d%%", i)

		// Flush the output if necessary (optional in most cases)
		// fmt.Flush() (if using a custom writer)

		// Simulate some work
		time.Sleep(50 * time.Millisecond)
	}

	// Print a newline to properly end the output
	fmt.Println()
}
