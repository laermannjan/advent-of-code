package main

import (
	"fmt"
	"lj/utils"
	"os"
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

func main() {
	input := utils.NewStdinInput()
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
		// fmt.Println(line, direction, length)
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
		fmt.Println(coords[i-1], "->", coords[i])

		p0 := shift(coords[i-1], min_p)
		p1 := shift(coords[i], min_p)
		this_area := (p0.row + p1.row) * (p0.col - p1.col)
		fmt.Println("this area", this_area)
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

	fmt.Fprintln(os.Stderr, i+circumference)
}
