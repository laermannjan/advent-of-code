package main

import (
	"fmt"
	"lj/utils"
	"os"
	"regexp"
)

type Position struct {
	row int
	col int
}

func shift(p Position, origin Position) Position {
	return Position{row: p.row - origin.row, col: p.col - origin.row}
}

func main() {
	input := utils.NewStdinInput()
	re := regexp.MustCompile(`(\w) (\d+) \(.(\w+)\)`)

	coords := map[Position]bool{}
	cur := Position{}
	min_p := Position{}
	max_p := Position{}

	for line := range input.Lines() {
		match := re.FindStringSubmatch(line)
		direction := match[1]
		length := utils.Atoi(match[2])
		fmt.Println(line, direction, length)
		// color := match[3]

		for i := 0; i < length; i++ {
			switch direction {
			case "U":
				cur.row--
				coords[cur] = true
				min_p.row = min(min_p.row, cur.row)
				max_p.row = max(max_p.row, cur.row)
			case "D":
				cur.row++
				coords[cur] = true
				min_p.row = min(min_p.row, cur.row)
				max_p.row = max(max_p.row, cur.row)
			case "L":
				cur.col--
				coords[cur] = true
				min_p.col = min(min_p.col, cur.col)
				max_p.col = max(max_p.col, cur.col)
			case "R":
				cur.col++
				coords[cur] = true
				min_p.col = min(min_p.col, cur.col)
				max_p.col = max(max_p.col, cur.col)
			default:
				fmt.Println("default")
			}
		}

	}
	fmt.Println("lagoon edge length", len(coords))
	pad := 1
	viz_pad := 2

	seed := Position{row: min_p.row - pad, col: min_p.col - pad}
	stack := []Position{}
	stack = append(stack, seed)
	seen := map[Position]bool{}
	for len(stack) > 0 {
		this := stack[0]
		stack = stack[1:]
		if !(min_p.row-1 <= this.row && this.row <= max_p.row+1) || !(min_p.col-1 <= this.col && this.col <= max_p.col+1) {
			continue
		}
		if seen[this] == true {
			continue
		}
		if coords[this] == true {
			continue
		}
		seen[this] = true
		stack = append(stack, Position{row: this.row, col: this.col - 1})
		stack = append(stack, Position{row: this.row, col: this.col + 1})
		stack = append(stack, Position{row: this.row - 1, col: this.col})
		stack = append(stack, Position{row: this.row + 1, col: this.col})
	}

	for r := min_p.row - viz_pad; r <= max_p.row+viz_pad; r++ {
		line := ""
		for c := min_p.col - 2; c <= max_p.col+2; c++ {
			if coords[Position{r, c}] == true {
				line += fmt.Sprint("#")
			} else if seen[Position{r, c}] == true {
				line += fmt.Sprint("-")
			} else {
				line += fmt.Sprint(".")
			}
		}
		fmt.Println(line)
	}
	fmt.Println("outside", len(seen))
	fmt.Fprintln(os.Stderr, ((max_p.row+pad)-(min_p.row-pad)+1)*((max_p.col+pad)-(min_p.col-pad)+1)-len(seen))
}
