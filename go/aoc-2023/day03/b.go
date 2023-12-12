package day03

import (
	"fmt"
	"log"
	"slices"
	"unicode"

	"github.com/laermannjan/advent-of-code/go/utils"
	"github.com/spf13/cobra"
)

func BCmd() *cobra.Command {
	part := "b"
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

			fmt.Printf("Answer: %d\n", partB(input))
		},
	}
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
