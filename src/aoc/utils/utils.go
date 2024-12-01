package utils

import (
	"log"
	"strconv"
)

func ToInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
