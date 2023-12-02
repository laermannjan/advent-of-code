package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
				input = utils.FromExampleFile(2023, 1, part)
			} else {
				input = utils.FromInputFile(2023, 1)
			}

			fmt.Printf("Answer: %d\n", partB(input))
		},
	}
}

func partB(input utils.Input) int {
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0
	for line := range input.Lines() {
		runes := []rune(line)

		var firstDigit int
		var secondDigit int

	left_letter_loop:
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit, _ = strconv.Atoi(string(runes[i]))
				break
			} else {
				for d, name := range digits {
					if strings.HasPrefix(string(runes[i:]), name) {
						firstDigit = d
						break left_letter_loop
					}
				}
			}
		}

	right_letter_loop:
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				secondDigit, _ = strconv.Atoi(string(runes[i]))
				break
			} else {
				for d, name := range digits {
					if strings.HasSuffix(string(runes[:i+1]), name) {
						secondDigit = d
						break right_letter_loop
					}
				}
			}
		}

		// log.Println(line)
		// log.Println(firstDigit, secondDigit)
		sum += firstDigit*10 + secondDigit
	}
	return sum
}
