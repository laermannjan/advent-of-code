package main

import (
	"fmt"
	"lj/utils"
	"os"
	"slices"
	"unicode"
)

type Pair struct {
	x int
	y int
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func getDirections(lines []string, x, y int) []Pair {
	all_dirs := []Pair{{1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}, {-1, -1}, {-1, 0}, {-1, 1}}
	possible_dirs := []Pair{}
	for _, dir := range all_dirs {
		check_y := y + dir.y
		if (0 > check_y) || (check_y >= len(lines)) {
			continue
		}
		check_line := lines[y+dir.y]
		check_x := x + dir.x
		if (0 > check_x) || (check_x >= len(check_line)) {
			continue
		}
		possible_dirs = append(possible_dirs, dir)
	}
	return possible_dirs
}

func main() {
	input := utils.NewStdinInput()

	sum := 0

	lines := input.LineSlice()
	stars := make(map[Pair][]int)
	for y, line := range lines {
		num := ""
		isPartNumber := false
		adjacentStars := []Pair{}
		for x, char := range line {
			fmt.Printf("check x=%d y=%d, num=%s\n", x, y, num)
			if unicode.IsDigit(char) {
				num += string(char)
				for _, dir := range getDirections(lines, x, y) {
					check_x := x + dir.x
					check_y := y + dir.y
					check_char := []rune(lines[check_y])[check_x]
					if !isPartNumber && isSymbol(check_char) {
						isPartNumber = true
					}
					if check_char == '*' {
						star_pos := Pair{check_x, check_y}
						if !slices.Contains(adjacentStars, star_pos) {
							adjacentStars = append(adjacentStars, star_pos)
						}
					}

				}
			}
			fmt.Println("adjacent stars:", adjacentStars)

			if !unicode.IsDigit(char) || x == len(line)-1 {
				for _, star := range adjacentStars {
					stars[star] = append(stars[star], utils.Atoi(num))
					fmt.Println(stars)
				}
				num = ""
				adjacentStars = []Pair{}

			}
		}
	}

	for _, parts := range stars {
		gear_ratio := 1
		if len(parts) == 2 {
			for _, part := range parts {
				gear_ratio *= part
			}
			sum += gear_ratio

		}
	}
	fmt.Fprintln(os.Stderr, sum)
}
