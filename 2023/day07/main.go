package main

import (
	"aoc-go/utils"
	"log"
	"slices"
	"strings"
)

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

func (h hand) get_type(joker bool) (typ int) {
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

	if joker {
		if joker_count, exists := card_counts['J']; exists {
			switch typ {
			case high_card:
				typ = one_pair
			case one_pair:
				typ = three_kind
			case two_pair:
				if joker_count == 1 {
					typ = full_house
				} else {
					typ = four_kind
				}
			case three_kind:
				typ = four_kind
			case full_house:
				typ = five_kind
			case four_kind:
				typ = five_kind
			}
		}

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

var cardValuesJoker = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func part1(input utils.Input) interface{} {
	hands := []hand{}
	for line := range input.Lines() {
		fields := strings.Fields(line)
		h := hand{cards: fields[0], bid: utils.Atoi(fields[1])}
		hands = append(hands, h)
		log.Println("hand", h, h.get_type(false))

	}
	slices.SortFunc(hands, func(a, b hand) int {
		typ_diff := a.get_type(false) - b.get_type(false)
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

func part2(input utils.Input) interface{} {
	hands := []hand{}
	for line := range input.Lines() {
		fields := strings.Fields(line)
		h := hand{cards: fields[0], bid: utils.Atoi(fields[1])}
		hands = append(hands, h)
		log.Println("hand", h, h.get_type(true))

	}
	slices.SortFunc(hands, func(a, b hand) int {
		typ_diff := a.get_type(true) - b.get_type(true)
		if typ_diff != 0 {
			return typ_diff
		}
		cards_a := []rune(a.cards)
		cards_b := []rune(b.cards)
		for i := 0; i < len(cards_a); i++ {

			if cards_a[i] != cards_b[i] {
				cmp_val := cardValuesJoker[cards_a[i]] - cardValuesJoker[cards_b[i]]
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

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
