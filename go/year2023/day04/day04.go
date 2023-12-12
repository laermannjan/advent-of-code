package day04

import (
	"fmt"
	"log"
	"reflect"
	"slices"
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
	total_score := 0
	for line := range input.Lines() {
		numbers := strings.Split(strings.Split(line, ":")[1], "|")
		winners := strings.Fields(numbers[0])
		chosen := strings.Fields(numbers[1])

		score := 1
		log.Print(winners, chosen)
		for _, n := range chosen {
			if slices.Contains(winners, n) {
				score <<= 1
				log.Println("score(shift):", score, n)

			}
		}

		log.Println("score:", score)
		// we started with score=1 instead of 0 in order to bitshift, so shift back
		total_score += score >> 1
	}

	return total_score
}

func partB(input utils.Input) int {
	card_count := map[int]int{}

	for line := range input.Lines() {
		schema := strings.Split(line, ":")
		card := utils.Atoi(strings.Fields(schema[0])[1])
		numbers := strings.Split(schema[1], "|")
		winners := strings.Fields(numbers[0])
		chosen := strings.Fields(numbers[1])

		card_count[card] += 1
		log.Println("Card:", card, winners, chosen, "copies:", card_count[card])

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
