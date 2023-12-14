package day07

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

const (
	high_card  = iota
	one_pair   = iota
	two_pair   = iota
	three_kind = iota
	full_house = iota
	four_kind  = iota
	five_kind  = iota
)

type hand struct {
	cards string
	bid   int
}

func (h hand) get_type() (typ int) {
	cards := []rune(h.cards)
	card_counts := map[rune]int{}

	for _, card := range cards {
		card_counts[card]++
	}

	if len(card_counts) == 1 {
		typ = five_kind
	} else if len(card_counts) == 2 {
		if card_counts[cards[0]] == 2 || card_counts[cards[0]] == 3 {
			typ = full_house
		} else {
			typ = four_kind
		}
	} else if len(card_counts) == 3 {
		if card_counts[cards[0]] == 3 || card_counts[cards[1]] == 3 || card_counts[cards[2]] == 3 {
			typ = three_kind
		} else {
			typ = two_pair
		}
	} else if len(card_counts) == 4 {
		typ = one_pair
	} else if len(card_counts) == 5 {
		typ = high_card
	}

	return
}

var cardValues = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func partA(input utils.Input) int {
	hands := []hand{}
	for line := range input.Lines() {
		fields := strings.Fields(line)
		h := hand{cards: fields[0], bid: utils.Atoi(fields[1])}
		hands = append(hands, h)
		log.Println("hand", h, h.get_type())

	}
	slices.SortFunc(hands, func(a, b hand) int {
		typ_diff := a.get_type() - b.get_type()
		if typ_diff != 0 {
			return typ_diff
		}
		cards_a := []rune(a.cards)
		cards_b := []rune(b.cards)
		for i := 0; i < len(cards_a); i++ {

			if cards_a[i] != cards_b[i] {
				cmp_val := cardValues[cards_a[i]] - cardValues[cards_b[i]]
				return cmp_val
			}
		}
		return 0

	})

	log.Println("\n", hands)

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * hand.bid
	}
	return winnings
}

func partB(input utils.Input) int {
	return 0
}
