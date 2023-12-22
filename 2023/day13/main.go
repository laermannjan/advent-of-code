package main

import (
	"aoc-go/utils"
	"log"
	"strings"
)

func transpose(pattern []string) []string {
	var transposed []string
	for col := range pattern[0] {
		var t_row []byte
		for row := range pattern {
			t_row = append(t_row, pattern[row][col])
		}
		transposed = append(transposed, string(t_row))
	}
	return transposed
}

func equal(a, b string) (eq bool, withSmudge bool) {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if !withSmudge {
				withSmudge = true
				continue
			} else {
				return false, false
			}
		}
	}
	return true, withSmudge
}

func findPointOfReflection(pattern []string, allowSmudge bool) int {
row_loop:
	for r := 1; r < len(pattern); r++ {
		smudge_used := false
		for i := 0; 0 <= r-1-i && r+i < len(pattern); i++ {
			eq, withSmudge := equal(pattern[r-1-i], pattern[r+i])

			log.Printf("comparing\n%v\n%v\neq=%v, withSmudge=%v, smudge_used=%v, allow_smudge=%v", strings.Join(pattern[r-1-i:r], "\n"), strings.Join(pattern[r:r+i+1], "\n"), eq, withSmudge, smudge_used, allowSmudge)
			if !(eq && (!withSmudge || (allowSmudge && !smudge_used))) {
				log.Println("NO MATCH, continue\n")
				continue row_loop
			}
			if eq && withSmudge && !smudge_used {
				smudge_used = true
			}
		}
		if !allowSmudge || smudge_used {
			return r
		}
		log.Printf("no smudge :( - try next")
	}
	return 0
}

func part1(input utils.Input) (answer interface{}) {
	total := 0
	for pattern := range input.Sections() {
		log.Printf("\n%v", pattern)
		p := strings.Split(pattern, "\n")
		rows_above := findPointOfReflection(p, false)
		if rows_above > 0 {
			log.Println("point of reflection (rows):", rows_above)
		}
		cols_before := findPointOfReflection(transpose(p), false)
		if cols_before > 0 {
			log.Println("point of reflection (cols):", cols_before)
		}

		total += 100*rows_above + cols_before
	}
	return total
}

func part2(input utils.Input) (answer interface{}) {
	total := 0
	for pattern := range input.Sections() {
		log.Printf("\n%v", pattern)
		p := strings.Split(pattern, "\n")
		rows_above := findPointOfReflection(p, true)
		if rows_above > 0 {
			log.Println("point of reflection (rows):", rows_above)
		}
		cols_before := findPointOfReflection(transpose(p), true)
		if cols_before > 0 {
			log.Println("point of reflection (cols):", cols_before)
		}

		total += 100*rows_above + cols_before
	}
	return total
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
