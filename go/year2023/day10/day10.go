package day10

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"slices"
	"strings"

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

var options = map[byte][2]string{
	'L': {"north", "east"},
	'|': {"north", "south"},
	'J': {"north", "west"},
	'F': {"east", "south"},
	'-': {"east", "west"},
	'7': {"south", "west"},
}

var inverse = map[string]string{
	"north": "south",
	"south": "north",
	"west":  "east",
	"east":  "west",
}

var enters = map[string][]byte{
	"north": {'|', '7', 'F'},
	"south": {'|', 'J', 'L'},
	"west":  {'-', 'L', 'F'},
	"east":  {'-', 'J', '7'},
}

type position struct {
	row int
	col int
}

func (from position) move(maze []string, dir string) (to position, err error) {
	to = from
	switch dir {
	case "north":
		to.row -= 1
		if to.row < 0 {
			err = errors.New("already on north edge")
		}
	case "south":
		to.row += 1
		if to.row >= len(maze) {
			err = errors.New("already on south edge")
		}
	case "west":
		to.col -= 1
		if to.col < 0 {
			err = errors.New("already on west edge")
		}
	case "east":
		to.col += 1
		if to.col > len(maze[0]) {
			err = errors.New("already on east edge")
		}
	}
	return
}

func (p position) symbol(grid []string) (symbol byte) {
	symbol = grid[p.row][p.col]
	return
}

func partA(input utils.Input) int {
	maze := input.LineSlice()
	var s position
	for r := range maze {
		for c, ch := range maze[r] {
			if ch == 'S' {
				s = position{row: r, col: c}
			}
		}
	}

	loop := map[position]bool{s: true}
	current := s
	log.Println("starting:", current)
	prev_move := ""
	for {
		if current == s {
			for move, valid_symbols := range enters {
				next, err := current.move(maze, move)
				if err != nil {
					continue
				}
				if slices.Contains(valid_symbols, next.symbol(maze)) {
					current = next
					prev_move = move
					loop[current] = true
					log.Println("move:", prev_move)
					break
				}
			}
		}

		opts := options[current.symbol(maze)]
		if inverse[opts[0]] == prev_move {
			prev_move = opts[1]
		} else {
			prev_move = opts[0]

		}
		log.Println("move:", prev_move)
		current, _ = current.move(maze, prev_move)
		loop[current] = true
		if current == s {
			break
		}
	}

	return len(loop) / 2
}

// a point within the loop needs to cross the loop an uneven number of times
// and a point outside will not cross it at all or cross it an even number of times
// we can just go from left to right per line and count how many times we cross the pipe
// all free cells '.' where we crossed the pipe an uneven amount of time before will be within the loop
// special case is where we have constructs like 'F---7' (down-turn, down-turn), which we don't cross, or 'F---J' (opposite turns) which we need to cross once.

func partB(input utils.Input) int {
	maze := input.LineSlice()
	var s position
	for r := range maze {
		for c, ch := range maze[r] {
			if ch == 'S' {
				s = position{row: r, col: c}
			}
		}
	}

	loop := map[position]bool{s: true}
	current := s
	log.Println("starting:", current)
	prev_move := ""

	for {
		if current == s {
			for move, valid_symbols := range enters {
				next, err := current.move(maze, move)
				if err != nil {
					continue
				}
				if slices.Contains(valid_symbols, next.symbol(maze)) {
					current = next
					prev_move = move
					loop[current] = true
					log.Println("move:", prev_move)
					break
				}
			}
		}

		opts := options[current.symbol(maze)]
		if inverse[opts[0]] == prev_move {
			prev_move = opts[1]
		} else {
			prev_move = opts[0]

		}
		log.Println("move:", prev_move)
		current, _ = current.move(maze, prev_move)
		loop[current] = true
		if current == s {
			break
		}
	}

	// fmt.Println(strings.Join(maze, "\n"))
	for r := range maze {
		for c := range maze[0] {
			if _, ok := loop[position{row: r, col: c}]; !ok {
				maze[r] = maze[r][:c] + "." + maze[r][c+1:]
			}
		}
	}

	// find pipe for 'S'
	s_dirs := [2]string{}
	si := 0
	for _, dir := range []string{"north", "east", "south", "west"} {
		if pos, ok := s.move(maze, dir); ok == nil && slices.Contains(enters[dir], pos.symbol(maze)) {
			log.Println("dir", dir, "pos", pos)
			s_dirs[si] = dir
			si++
		}
	}
	for symbol, opts := range options {
		if s_dirs == opts {
			fmt.Println("converting", string(s.symbol(maze)), "to", string(symbol))
			maze[s.row] = maze[s.row][:s.col] + string(symbol) + maze[s.row][s.col+1:]
			break
		}
	}

	fmt.Println(strings.Join(maze, "\n"))

	n_within := 0
	for r, row := range maze {
		var last rune
		crossed := false
		for c, ch := range row {
			if ch == '|' {
				crossed = !crossed
			} else if ch == 'L' || ch == 'F' {
				last = ch
			} else if ch == 'J' {
				if last == 'F' {
					crossed = !crossed
				}
				last = 0
			} else if ch == '7' {
				if last == 'L' {
					crossed = !crossed
				}
				last = 0
			} else if ch == '.' && crossed {
				n_within++
				log.Println(r, c, "is within")

			}
		}
	}

	return n_within
}
