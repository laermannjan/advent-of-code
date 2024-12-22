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

	card_count := map[int]int{}

	for line := range input.Lines() {
		schema := strings.Split(line, ":")
		card := utils.Atoi(strings.Fields(schema[0])[1])
		numbers := strings.Split(schema[1], "|")
		winners := strings.Fields(numbers[0])
		chosen := strings.Fields(numbers[1])

		card_count[card] += 1
		fmt.Println("Card:", card, winners, chosen, "copies:", card_count[card])

		matches := 0
		for _, n := range chosen {
			if slices.Contains(winners, n) {
				matches += 1
				card_count[card+matches] += card_count[card]
			}
		}

	}

	total_cards := 0
	for _, counts := range card_count {
		total_cards += counts
	}

	fmt.Fprintln(os.Stderr, total_cards)
}
