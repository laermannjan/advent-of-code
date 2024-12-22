package main

import (
	"fmt"
	"lj/utils"
	"os"
	"strings"
)

func hash(input string) int {
	cur := 0
	for _, ch := range input {
		cur += int(ch)
		cur *= 17
		cur %= 256
	}
	return cur
}

func main() {
	input := utils.NewStdinInput()
	codes := strings.Split(input.LineSlice()[0], ",")
	total := 0
	for _, code := range codes {
		total += hash(code)
	}
	fmt.Fprintln(os.Stderr, total)
}
