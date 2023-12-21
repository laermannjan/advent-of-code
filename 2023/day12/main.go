package main

import (
	"aoc-go/utils"
	"log"
	"strings"
)

func combinations(cfg string) (combos []string) {
	if cfg == "" {
		return []string{""}
	}
	if cfg[0] == '?' {
		for _, combo := range combinations(cfg[1:]) {
			combos = append(combos, "#"+combo)
			combos = append(combos, "."+combo)
		}
	} else {
		for _, combo := range combinations(cfg[1:]) {
			combos = append(combos, string(cfg[0])+combo)
		}
	}
	return
}

func part1(input utils.Input) (answer interface{}) {
	total_combos := 0
	for _, line := range input.LineSlice() {
		parts := strings.Split(line, " ")
		config := parts[0]
		var groups []int
		for _, g := range strings.Split(parts[1], ",") {
			groups = append(groups, utils.Atoi(g))
		}

		log.Println()
		log.Println(config, groups)
	combo_loop:
		for _, combo := range combinations(config) {
			var combo_groups []int
			for _, g := range strings.Split(combo, ".") {
				if len(g) > 0 {
					combo_groups = append(combo_groups, len(g))
				}
			}

			if len(combo_groups) == len(groups) {
				for i := range combo_groups {
					if combo_groups[i] != groups[i] {
						continue combo_loop
					}
				}
				// log.Println("combo:", combo, combo_groups)
				total_combos++
			}
		}

	}
	return total_combos
}

func part2(input utils.Input) (answer interface{}) {
	return
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
