package main

import (
	"aoc-go/utils"
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

func part1(input utils.Input) (answer interface{}) {
	codes := strings.Split(input.LineSlice()[0], ",")
	total := 0
	for _, code := range codes {
		total += hash(code)
	}
	return total
}

func part2(input utils.Input) (answer interface{}) {
	return
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
