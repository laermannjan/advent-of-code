package day01

import (
	"fmt"
	"log"
	"strconv"
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
				input = utils.FromExampleFile(2023, 1, part)
			} else {
				input = utils.FromInputFile(2023, 1)
			}

			fmt.Printf("Answer: %d\n", partA(input))
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
