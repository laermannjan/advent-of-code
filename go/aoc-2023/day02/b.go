package day02

import (
	"fmt"
	"log"
	"strings"

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
				input = utils.FromExampleFile(2023, 2, part)
			} else {
				input = utils.FromInputFile(2023, 2)
			}

			fmt.Printf("Answer: %d\n", partB(input))
		},
	}
}

func partB(input utils.Input) int {
	sum := 0

	for line := range input.Lines() {
		needed_cubes := map[string]int{}
		parts := strings.Split(line, ":")
		game := strings.Split(parts[1], ";")

		for _, set := range game {
			cubes := strings.Split(set, ",")
			for _, color := range cubes {
				v := strings.Fields(color)
				count := utils.Atoi(v[0])
				log.Println(v[1], count, needed_cubes[v[1]])

				needed_cubes[v[1]] = max(needed_cubes[v[1]], count)
			}
		}

		product := 1
		for _, v := range needed_cubes {
			product *= v
		}

		sum += product
	}
	return sum
}
