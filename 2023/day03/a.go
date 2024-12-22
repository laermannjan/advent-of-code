package main

import (
	"fmt"
	"lj/utils"
	"os"
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
	for y, line := range lines {
		num := ""
		isPartNumber := false
		for x, char := range line {
			if unicode.IsDigit(char) {
				num += string(char)
				if !isPartNumber {
					for _, dir := range getDirections(lines, x, y) {
						if isSymbol([]rune(lines[y+dir.y])[x+dir.x]) {
							isPartNumber = true
							break
						}
					}
				}
			}

			if !unicode.IsDigit(char) || x == len(line)-1 {
				if isPartNumber {
					sum += utils.Atoi(num)
					fmt.Println(sum, num)
					isPartNumber = false
				}
				num = ""
			}
		}
	}
	fmt.Fprintln(os.Stderr, sum)
}
