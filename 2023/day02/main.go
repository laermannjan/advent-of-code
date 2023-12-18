package main

import (
	"aoc-go/utils"
	"log"
	"strings"
)

var max_cubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func part1(input utils.Input) interface{} {
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

func part2(input utils.Input) interface{} {
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

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
