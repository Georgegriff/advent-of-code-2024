package readfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type OnReadLine func(line string) error

func Open(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func ReadLine(f *os.File, callback OnReadLine) error {

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if err := callback(line); err != nil {
			return fmt.Errorf("callback error: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return nil

}

func ToInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
