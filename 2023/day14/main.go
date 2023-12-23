package main

import (
	"aoc-go/utils"
	"log"
)

func transpose(pattern []string) []string {
	var rune_pattern [][]rune
	for _, row := range pattern {
		rune_pattern = append(rune_pattern, []rune(row))
	}

	var transposed []string
	for col := range rune_pattern[0] {
		var t_row []rune
		for row := range rune_pattern {
			t_row = append(t_row, rune_pattern[row][col])
		}
		transposed = append(transposed, string(t_row))
	}
	return transposed
}

func part1(input utils.Input) (answer interface{}) {
	dish := input.LineSlice()

	load := 0

	t_dish := transpose(dish)
	var shifted []string
	for _, row := range t_dish {
		skipped := 0
		var shifted_row []rune
		for _, stone := range row {
			switch stone {
			case '.':
				skipped++
			case 'O':
				load += len(row) - len(shifted_row)
				shifted_row = append(shifted_row, stone)
			case '#':
				for i := 0; i < skipped; i++ {
					shifted_row = append(shifted_row, '.')
				}
				shifted_row = append(shifted_row, '#')
				skipped = 0
			}
		}
		for i := 0; i < skipped; i++ {
			shifted_row = append(shifted_row, '.')
		}
		shifted = append(shifted, string(shifted_row))
	}
	shifted = transpose(shifted)
	log.Println("original\tshifted north")
	for i := 0; i < len(dish); i++ {
		log.Println(dish[i], "\t", shifted[i])
	}

	return load
}

func part2(input utils.Input) (answer interface{}) {
	return
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
