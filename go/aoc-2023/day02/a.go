package day02

import (
	"fmt"
	"log"
	"strings"

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
				input = utils.FromExampleFile(2023, 2, part)
			} else {
				input = utils.FromInputFile(2023, 2)
			}

			fmt.Printf("Answer: %d\n", partA(input))
		},
	}
}

var max_cubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func partA(input utils.Input) int {
	sum := 0

game_loop:
	for line := range input.Lines() {
		parts := strings.Split(line, ":")
		game_id := utils.Atoi(strings.TrimPrefix(parts[0], "Game "))

		game := strings.Split(parts[1], ";")

		for _, set := range game {
			cubes := strings.Split(set, ",")
			for _, color := range cubes {
				v := strings.Fields(color)
				count := utils.Atoi(v[0])
				log.Println(v[1], count, max_cubes[v[1]])
				if count > max_cubes[v[1]] {
					continue game_loop
				}
			}
		}
		sum += game_id
	}
	return sum
}
