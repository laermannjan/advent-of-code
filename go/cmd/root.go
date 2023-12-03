package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-aoc",
		Short: "Advent of Code Solutions",
		Long:  "Golang implementations various years of the Advent of Code calendar",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if !debug {
				log.SetOutput(io.Discard)
			}
		},
	}

	debug bool
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug logging")
}
