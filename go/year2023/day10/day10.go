package day10

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

type node struct {
	symbol rune
	row    int
	col    int
	dist   int
}

var options = map[rune][2]string{
	'L': {"north", "east"},
	'7': {"west", "south"},
	'J': {"north", "west"},
	'F': {"south", "east"},
	'|': {"north", "south"},
	'-': {"west", "east"},
}

var inverse = map[string]string{
	"north": "south",
	"south": "north",
	"west":  "east",
	"east":  "west",
}

func walk(pipe rune, coming_from string) string {
	opts := options[pipe]
	if opts[0] == coming_from {
		return opts[1]
	} else {
		return opts[0]
	}
}

func partA(input utils.Input) int {
	maze := [][]node{}
	var s *node
	for i, line := range input.LineSlice() {
		row := []node{}
		for j, r := range line {
			n := node{symbol: r, col: j, row: i}
			row = append(row, n)
			if s == nil && r == 'S' {
				s = &n
				log.Printf("s found: %#v", n)
			}

		}
		maze = append(maze, row)
	}

	cur_node := &*s

	init_dirs := []string{}
	for _, dir := range []string{"north", "south", "west", "east"} {
		var pipe rune
		switch dir {
		case "north":
			if 0 <= cur_node.row-1 {
				pipe = maze[cur_node.row-1][cur_node.col].symbol
			}
		case "south":
			if cur_node.row+1 < len(maze) {
				pipe = maze[cur_node.row+1][cur_node.col].symbol
			}
		case "west":
			if 0 <= cur_node.col-1 {
				pipe = maze[cur_node.row][cur_node.col-1].symbol
			}
		case "east":
			if cur_node.col-1 < len(maze[cur_node.row]) {
				pipe = maze[cur_node.row][cur_node.col+1].symbol
			}
		}
		opts := options[pipe]
		if opts[0] == inverse[dir] || opts[1] == inverse[dir] {
			init_dirs = append(init_dirs, dir)
		}
	}
	log.Println("init dirs:", init_dirs)

	max_dist := 0
	for _, init_dir := range init_dirs {
		going_to := init_dir
		log.Printf("at %+v (%v) - going to [init] %s", cur_node, string(cur_node.symbol), going_to)
		dist := 0
		max_dist = 0
		for {
			switch going_to {
			case "north":
				cur_node = &maze[cur_node.row-1][cur_node.col]
			case "south":
				cur_node = &maze[cur_node.row+1][cur_node.col]
			case "west":
				cur_node = &maze[cur_node.row][cur_node.col-1]
			case "east":
				cur_node = &maze[cur_node.row][cur_node.col+1]
			}

			if cur_node.symbol == 'S' {
				break
			}
			dist++
			// log.Println("dists", dist, cur_node.dist)
			if cur_node.dist == 0 || cur_node.dist > dist {
				cur_node.dist = dist
			}
			max_dist = max(cur_node.dist, max_dist)
			going_to = walk(cur_node.symbol, inverse[going_to])
			log.Printf("at %+v (%v) - going to %s", cur_node, string(cur_node.symbol), going_to)
		}
	}
	return max_dist
}

func partB(input utils.Input) int {
	return 0
}
