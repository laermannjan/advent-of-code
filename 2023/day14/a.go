package main

import (
	"fmt"
	"lj/utils"
	"os"
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

func main() {
	input := utils.NewStdinInput()
	var dish [][]rune
	for _, row := range input.LineSlice() {
		dish = append(dish, []rune(row))
	}

	shifted := tilt(dish)
	load := computeLoad(shifted)

	fmt.Println("original\tshifted north")
	for i := 0; i < len(dish); i++ {
		fmt.Println(string(dish[i]), "\t", string(shifted[i]))
	}

	fmt.Fprintln(os.Stderr, load)
}
