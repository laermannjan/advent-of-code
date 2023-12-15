package day09

import (
	"fmt"
	"log"
	"reflect"

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

func partB(input utils.Input) int {
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
