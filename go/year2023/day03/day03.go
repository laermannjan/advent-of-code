package day03

import (
	"fmt"
	"log"
	"reflect"
	"slices"
	"unicode"

	"github.com/laermannjan/advent-of-code/go/utils"
	"github.com/spf13/cobra"
)

func ACmd() *cobra.Command {
	part := "a"
	return &cobra.Command{
		Use:   part,
		Short: "Part " + part,
		Run: func(cmd *cobra.Command, _ []string) {
			type PkgMark struct{}
			year, day := utils.GetYearDay(reflect.TypeOf(PkgMark{}).PkgPath())

			solveExample, err := cmd.Flags().GetBool("example")
			if err != nil {
				log.Fatal(err)
			}

			var input utils.Input
			if solveExample {
				input = utils.FromExampleFile(year, day, part)
			} else {
				input = utils.FromInputFile(year, day)
			}

			fmt.Printf("Answer: %d\n", partA(input))
		},
	}
}

func BCmd() *cobra.Command {
	part := "b"
	return &cobra.Command{
		Use:   part,
		Short: "Part " + part,
		Run: func(cmd *cobra.Command, _ []string) {
			type PkgMark struct{}
			year, day := utils.GetYearDay(reflect.TypeOf(PkgMark{}).PkgPath())

			solveExample, err := cmd.Flags().GetBool("example")
			if err != nil {
				log.Fatal(err)
			}

			var input utils.Input
			if solveExample {
				input = utils.FromExampleFile(year, day, part)
			} else {
				input = utils.FromInputFile(year, day)
			}

			fmt.Printf("Answer: %d\n", partB(input))
		},
	}
}

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

func partA(input utils.Input) int {
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
					log.Println(sum, num)
					isPartNumber = false
				}
				num = ""
			}
		}
	}
	return sum
}

func partB(input utils.Input) int {
	sum := 0

	lines := input.LineSlice()
	stars := make(map[Pair][]int)
	for y, line := range lines {
		num := ""
		isPartNumber := false
		adjacentStars := []Pair{}
		for x, char := range line {
			log.Printf("check x=%d y=%d, num=%s\n", x, y, num)
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
			log.Println("adjacent stars:", adjacentStars)

			if !unicode.IsDigit(char) || x == len(line)-1 {
				for _, star := range adjacentStars {
					stars[star] = append(stars[star], utils.Atoi(num))
					log.Println(stars)
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
	return sum
}
