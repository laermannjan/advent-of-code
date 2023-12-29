package main

import (
	"aoc-go/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Position struct {
	row int
	col int
}

func shift(p Position, origin Position) Position {
	return Position{row: p.row - origin.row, col: p.col - origin.row}
}

func part1(input utils.Input) (answer interface{}) {
	re := regexp.MustCompile(`(\w) (\d+) \(.(\w+)\)`)

	coords := map[Position]bool{}
	cur := Position{}
	min_p := Position{}
	max_p := Position{}

	for line := range input.Lines() {
		match := re.FindStringSubmatch(line)
		direction := match[1]
		length := utils.Atoi(match[2])
		log.Println(line, direction, length)
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
				log.Println("default")
			}
		}

	}
	log.Println("lagoon edge length", len(coords))
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
		log.Println(line)
	}
	log.Println("outside", len(seen))
	return ((max_p.row+pad)-(min_p.row-pad)+1)*((max_p.col+pad)-(min_p.col-pad)+1) - len(seen)
}

func part2(input utils.Input) (answer interface{}) {
	re := regexp.MustCompile(`(\w) (\d+) \(.(\w+)\)`)

	coords := []Position{{0, 0}}
	cur := Position{}
	min_p := Position{}
	max_p := Position{}
	circumference := 0

	for line := range input.Lines() {
		match := re.FindStringSubmatch(line)
		// direction := match[1]
		// length := utils.Atoi(match[2])

		inst := []rune(match[3])

		length_, _ := strconv.ParseInt(string(inst[:len(inst)-1]), 16, 0)
		length := int(length_)
		direction := string(inst[len(inst)-1])
		// log.Println(line, direction, length)
		circumference += length

		switch direction {
		case "3": //"U":
			cur.row -= length
			coords = append(coords, cur)
			min_p.row = min(min_p.row, cur.row)
			max_p.row = max(max_p.row, cur.row)
		case "1": //"D":
			cur.row += length
			coords = append(coords, cur)
			min_p.row = min(min_p.row, cur.row)
			max_p.row = max(max_p.row, cur.row)
		case "2": //"L":
			cur.col -= length
			coords = append(coords, cur)
			min_p.col = min(min_p.col, cur.col)
			max_p.col = max(max_p.col, cur.col)
		case "0": //"R":
			cur.col += length
			coords = append(coords, cur)
			min_p.col = min(min_p.col, cur.col)
			max_p.col = max(max_p.col, cur.col)
		}

	}

	// Trapezoid fomular to compute area of ploygon
	// https://en.wikipedia.org/wiki/Shoelace_formula#Trapezoid_formula
	area := 0
	for i := 1; i < len(coords); i++ {
		log.Println(coords[i-1], "->", coords[i])

		p0 := shift(coords[i-1], min_p)
		p1 := shift(coords[i], min_p)
		this_area := (p0.row + p1.row) * (p0.col - p1.col)
		log.Println("this area", this_area)
		area += this_area
	}

	A := float64(area) / 2.0

	// since we're in discreet space, we miscalculate the inside are.
	// the polygon edge can be thought of as running through the middle of each cell
	// or each cubic meter that we dig.
	// The "half" outside the middle line throught, will be missing from the area
	// calculated above.
	// The problem are corner cells (where we turn), because these are missing 3/4 or 1/4 of the area
	// depending on whether they are inside or outside the polygon. Determining which type of corner
	// they are is tidious.
	// Instead, use Pick's theorem to commpute the number of strictly interior points of polygon and add
	// back the boundary
	// https://en.wikipedia.org/wiki/Pick%27s_theorem

	i := int(A - float64(circumference)/2 + 1)

	return i + circumference
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
