package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	solveCmd = &cobra.Command{
		Use:   "solve",
		Short: "Solve an AOC problem",
		Long:  "Solve an Advent of Code problem for a specific year, day and part. Either solves the example or the actual input",
	}
	Example bool
)

func printFlags(cmd *cobra.Command) {
	fmt.Println("Local Flags:")
	printFlagSet(cmd.LocalFlags())

	fmt.Println("Persistent Flags:")
	printFlagSet(cmd.PersistentFlags())
}

func printFlagSet(flagSet *pflag.FlagSet) {
	flagSet.VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("- %s: %s (default: %s)\n", flag.Name, flag.Usage, flag.DefValue)
	})
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&Example, "example", false, "use example input")
	rootCmd.AddCommand(solveCmd)
}
