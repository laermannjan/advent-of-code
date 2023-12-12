package day01

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
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

func partA(input utils.Input) int {
	sum := 0
	for line := range input.Lines() {
		runes := []rune(line)

		var v int
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				vv, _ := strconv.Atoi(string(runes[i]))
				v += vv * 10
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				vv, _ := strconv.Atoi(string(runes[i]))
				v += vv
				break
			}
		}

		// log.Println(v)
		sum += v
	}
	return sum
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
