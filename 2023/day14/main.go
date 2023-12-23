package main

import (
	"aoc-go/utils"
	"log"
)

func tilt(dish [][]rune) (shifted [][]rune, load int) {
	for c := 0; c < len(dish); c++ {
		skipped := 0
		var shifted_col []rune
		for r := range dish {
			stone := dish[r][c]
			switch stone {
			case '.':
				skipped++
			case 'O':
				load += len(dish) - len(shifted_col)
				shifted_col = append(shifted_col, stone)
			case '#':
				for i := 0; i < skipped; i++ {
					shifted_col = append(shifted_col, '.')
				}
				shifted_col = append(shifted_col, '#')
				skipped = 0
			}
		}
		for i := 0; i < skipped; i++ {
			shifted_col = append(shifted_col, '.')
		}

		shifted = append(shifted, shifted_col)
	}
	return
}

func part1(input utils.Input) (answer interface{}) {
	var dish [][]rune
	for _, row := range input.LineSlice() {
		dish = append(dish, []rune(row))
	}

	shifted, load := tilt(dish)

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
