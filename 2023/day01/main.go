package main

import (
	"aoc-go/utils"
	"strconv"
	"strings"
	"unicode"
)

func part1(input utils.Input) interface{} {
	sum := 0
	for line := range input.Lines() {
		runes := []rune(line)

		var v int
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				vv, _ := strconv.Atoi(string(runes[i]))
				v += vv * 10
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				vv, _ := strconv.Atoi(string(runes[i]))
				v += vv
				break
			}
		}

		// log.Println(v)
		sum += v
	}
	return sum
}

func part2(input utils.Input) interface{} {
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0
	for line := range input.Lines() {
		runes := []rune(line)

		var firstDigit int
		var secondDigit int

	left_letter_loop:
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit, _ = strconv.Atoi(string(runes[i]))
				break
			} else {
				for d, name := range digits {
					if strings.HasPrefix(string(runes[i:]), name) {
						firstDigit = d
						break left_letter_loop
					}
				}
			}
		}

	right_letter_loop:
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				secondDigit, _ = strconv.Atoi(string(runes[i]))
				break
			} else {
				for d, name := range digits {
					if strings.HasSuffix(string(runes[:i+1]), name) {
						secondDigit = d
						break right_letter_loop
					}
				}
			}
		}

		// log.Println(line)
		// log.Println(firstDigit, secondDigit)
		sum += firstDigit*10 + secondDigit
	}
	return sum
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
