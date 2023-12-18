package main

import (
	"aoc-go/utils"
	"log"
)

func part1(input utils.Input) interface{} {
	pred_sum := 0
	for line := range input.Lines() {
		history := utils.ParseInts(line)

		log.Println("history:", history)
		last_nums := []int{history[len(history)-1]}

		for {

			diffs := []int{}
			all_zero := true
			for i, val := range history[1:] {
				diff := val - history[i] // i is offset by 1
				diffs = append(diffs, diff)
				all_zero = all_zero && diff == 0
			}
			log.Println("diffs:", diffs)

			if all_zero {
				break
			}
			history = diffs
			last_nums = append(last_nums, diffs[len(diffs)-1])

		}

		pred := 0
		for i := len(last_nums) - 1; i >= 0; i-- {
			pred += last_nums[i]
			log.Println(i, "pred:", pred)

		}

		pred_sum += pred

	}
	return pred_sum
}

func part2(input utils.Input) interface{} {
	pred_sum := 0
	for line := range input.Lines() {
		history := utils.ParseInts(line)

		log.Println("history:", history)
		first_nums := []int{history[0]}

		for {

			diffs := []int{}
			all_zero := true
			for i, val := range history[1:] {
				diff := val - history[i] // i is offset by 1
				diffs = append(diffs, diff)
				all_zero = all_zero && diff == 0
			}
			log.Println("diffs:", diffs)

			if all_zero {
				break
			}
			history = diffs
			first_nums = append(first_nums, diffs[0])

		}

		pred := 0
		for i := len(first_nums) - 1; i >= 0; i-- {
			pred = first_nums[i] - pred
			log.Println(i, "pred:", pred)

		}

		pred_sum += pred

	}
	return pred_sum
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
