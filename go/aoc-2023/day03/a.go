package day03

import (
	"fmt"
	"log"
	// "strconv"
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
			solveExample, err := cmd.Flags().GetBool("example")
			if err != nil {
				log.Fatal(err)
			}

			var input utils.Input
			if solveExample {
				input = utils.FromExampleFile(2023, 3, part)
			} else {
				input = utils.FromInputFile(2023, 3)
			}

			fmt.Printf("Answer: %d\n", partA(input))
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
