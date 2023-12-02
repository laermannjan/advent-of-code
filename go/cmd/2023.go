/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laermannjan/advent-of-code/go/aoc-2023/day01"
	"github.com/spf13/cobra"
)

var year2023Cmd = &cobra.Command{
	Use:   "2023",
	Short: "AOC of 2023",
	Long:  "The Advent of Code puzzle implementations of 2023",
}

func makeDayCommand(day string, aCmd *cobra.Command, bCmd *cobra.Command) *cobra.Command {
	dayCmd := &cobra.Command{
		Use:   day,
		Short: "Solve the puzzle of " + day,
		Run: func(_ *cobra.Command, _ []string) {
			aCmd.Run(aCmd, []string{})
			bCmd.Run(bCmd, []string{})
		},
	}
	aCmd.Flags().Bool("example", false, "Solve the example input")
	bCmd.Flags().Bool("example", false, "Solve the example input")
	// dayCmd.PersistentFlags().Bool("example", false, "Solve the example input")

	dayCmd.AddCommand(aCmd)
	dayCmd.AddCommand(bCmd)
	return dayCmd
}

func init() {
	solveCmd.AddCommand(year2023Cmd)
	year2023Cmd.AddCommand(makeDayCommand("1", day01.ACmd(), day01.BCmd()))
}
