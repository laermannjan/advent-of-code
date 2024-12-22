package main

import (
	"aoc-go/utils"
	"fmt"
	"slices"
	"strings"
)

func part1(input utils.Input) interface{} {
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

	return total_score
}

func part2(input utils.Input) interface{} {
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

	return total_cards
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
