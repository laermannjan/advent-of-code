package main

import (
	"aoc-go/utils"
	"log"
	"strings"
)

func rotate(dish [][]rune) [][]rune {
	// rotate clockwise
	var rotated [][]rune

	for c := 0; c < len(dish[0]); c++ {
		var rot_row []rune
		for r := len(dish) - 1; r >= 0; r-- {
			rot_row = append(rot_row, dish[r][c])
		}
		rotated = append(rotated, rot_row)
	}
	return rotated
}

func computeLoad(dish [][]rune) (load int) {
	for r, row := range dish {
		for _, cell := range row {
			if cell == 'O' {
				load += len(dish) - r
			}
		}
	}
	return
}

func tilt(dish [][]rune) [][]rune {
	shifted := make([][]rune, len(dish))
	for i := range shifted {
		shifted[i] = make([]rune, len(dish[0]))
	}
	for c := 0; c < len(dish); c++ {
		skipped := 0
		shift_r := 0
		for r := range dish {
			stone := dish[r][c]
			switch stone {
			case '.':
				skipped++
			case 'O':
				shifted[shift_r][c] = stone
				shift_r++
			case '#':
				for i := 0; i < skipped; i++ {
					shifted[shift_r][c] = '.'
					shift_r++
				}
				shifted[shift_r][c] = '#'
				shift_r++
				skipped = 0
			}
		}

		for i := 0; i < skipped; i++ {
			shifted[shift_r][c] = '.'
			shift_r++
		}
	}
	return shifted
}

func part1(input utils.Input) (answer interface{}) {
	var dish [][]rune
	for _, row := range input.LineSlice() {
		dish = append(dish, []rune(row))
	}

	shifted := tilt(dish)
	load := computeLoad(shifted)

	log.Println("original\tshifted north")
	for i := 0; i < len(dish); i++ {
		log.Println(string(dish[i]), "\t", string(shifted[i]))
	}

	return load
}

func hash(dish [][]rune) string {
	var s []string
	for _, row := range dish {
		s = append(s, string(row))
	}
	return strings.Join(s, "\n")
}

func part2(input utils.Input) (answer interface{}) {
	target_cycle := 1_000_000_000
	var dish [][]rune
	for _, row := range input.LineSlice() {
		dish = append(dish, []rune(row))
	}

	var load int
	var shifted [][]rune

	seen := map[string]int{}

	jumped := false
	for cycle := 1; cycle <= target_cycle; cycle++ {
		for i := 0; i < 4; i++ {
			// old_dish := dish
			shifted = tilt(dish)
			dish = rotate(shifted)
			// log.Println("\noriginal\tshifted north\trotated")
			// for i := 0; i < len(dish); i++ {
			// 	log.Println(string(old_dish[i]), "\t", string(shifted[i]), "\t", string(dish[i]))
			// }

		}
		load = computeLoad(dish)
		log.Println("After", cycle, "cycle", "load", load)
		for i := 0; i < len(dish); i++ {
			log.Println(string(string(dish[i])))
		}
		if v, ok := seen[hash(shifted)]; ok && !jumped {
			length := cycle - v
			cycle = target_cycle - ((target_cycle - cycle) % length)
			jumped = true
		}
		seen[hash(shifted)] = cycle

	}
	load = computeLoad(dish)
	return load
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
