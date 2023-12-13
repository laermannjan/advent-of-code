/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/laermannjan/advent-of-code/go/year2023/day01"
	"github.com/laermannjan/advent-of-code/go/year2023/day02"
	"github.com/laermannjan/advent-of-code/go/year2023/day03"
	"github.com/laermannjan/advent-of-code/go/year2023/day04"
	"github.com/laermannjan/advent-of-code/go/year2023/day05"
	"github.com/laermannjan/advent-of-code/go/year2023/day06"
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
	dayCmd.AddCommand(aCmd)
	dayCmd.AddCommand(bCmd)
	return dayCmd
}

func init() {
	solveCmd.AddCommand(year2023Cmd)
	year2023Cmd.AddCommand(makeDayCommand("1", day01.ACmd(), day01.BCmd()))
	year2023Cmd.AddCommand(makeDayCommand("2", day02.ACmd(), day02.BCmd()))
	year2023Cmd.AddCommand(makeDayCommand("3", day03.ACmd(), day03.BCmd()))
	year2023Cmd.AddCommand(makeDayCommand("4", day04.ACmd(), day04.BCmd()))
	year2023Cmd.AddCommand(makeDayCommand("5", day05.ACmd(), day05.BCmd()))
	year2023Cmd.AddCommand(makeDayCommand("6", day06.ACmd(), day06.BCmd()))
}
