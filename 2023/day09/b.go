package main

import (
	"fmt"
	"lj/utils"
	"os"
)

func main() {
	input := utils.NewStdinInput()

	pred_sum := 0
	for line := range input.Lines() {
		history := utils.ParseInts(line)

		fmt.Println("history:", history)
		first_nums := []int{history[0]}

		for {

			diffs := []int{}
			all_zero := true
			for i, val := range history[1:] {
				diff := val - history[i] // i is offset by 1
				diffs = append(diffs, diff)
				all_zero = all_zero && diff == 0
			}
			fmt.Println("diffs:", diffs)

			if all_zero {
				break
			}
			history = diffs
			first_nums = append(first_nums, diffs[0])

		}

		pred := 0
		for i := len(first_nums) - 1; i >= 0; i-- {
			pred = first_nums[i] - pred
			fmt.Println(i, "pred:", pred)

		}

		pred_sum += pred

	}
	fmt.Fprintln(os.Stderr, pred_sum)
}
