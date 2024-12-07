package utils

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func ToInt(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func StringInSlice(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func MapToString(slice interface{}, format string) ([]string, error) {
	// Check if the input is a slice
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return nil, fmt.Errorf("provided input is not a slice")
	}

	// Create a result slice
	result := make([]string, val.Len())

	// Iterate through the slice and format each element
	for i := 0; i < val.Len(); i++ {
		result[i] = fmt.Sprintf(format, val.Index(i).Interface())
	}

	return result, nil
}
