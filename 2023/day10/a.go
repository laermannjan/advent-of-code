package main

import (
	"errors"
	"fmt"
	"lj/utils"
	"os"
	"slices"
)

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

func main() {
	input := utils.NewStdinInput()
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
	fmt.Println("starting:", current)
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
					fmt.Println("move:", prev_move)
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
		fmt.Println("move:", prev_move)
		current, _ = current.move(maze, prev_move)
		loop[current] = true
		if current == s {
			break
		}
	}

	fmt.Fprintln(os.Stderr, len(loop)/2)
}
