package main

import (
	"fmt"
	"lj/utils"
	"os"
	"slices"
	"strings"
)

func main() {
	input := utils.NewStdinInput()

	total_score := 0
	for line := range input.Lines() {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winners := strings.Fields(numbers[0])
		chosen := strings.Fields(numbers[1])

		score := 1
		fmt.Print(winners, chosen)
		for _, n := range chosen {
			if slices.Contains(winners, n) {
				score <<= 1
				fmt.Println("score(shift):", score, n)

			}
		}

		fmt.Println("score:", score)
		// we started with score=1 instead of 0 in order to bitshift, so shift back
		total_score += score >> 1
	}

	fmt.Fprintln(os.Stderr, total_score)
}
