package utils

import (
	"strconv"
	"strings"
)

func ParseInts(a string) (ints []int) {
	fields := strings.Fields(a)
	for _, f := range fields {
		ints = append(ints, Atoi(f))
	}
	return
}

func Atoi(a string) int {
	result, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	return result
}
