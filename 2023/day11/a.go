package main

import (
	"fmt"
	"lj/utils"
	"os"
)

type position struct {
	row int
	col int
}

func compute_galaxy_dists(input utils.Input, expansion_factor int) int {
	var galaxies []position

	for row, line := range input.LineSlice() {
		for col, ch := range line {
			if ch == '#' {
				galaxies = append(galaxies, position{row: row, col: col})
			}
		}
	}

	fmt.Println("found", len(galaxies), "galaxies")
	total_dist := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			row_lower := min(g1.row, g2.row)
			row_upper := max(g1.row, g2.row)
			col_lower := min(g1.col, g2.col)
			col_upper := max(g1.col, g2.col)

			rows := row_upper - row_lower
		row_loop:
			for r := row_lower + 1; r < row_upper; r++ {
				for _, g := range galaxies {
					if g.row == r {
						continue row_loop
					}
				}
				rows += expansion_factor - 1
			}

			cols := col_upper - col_lower
		col_loop:
			for r := col_lower + 1; r < col_upper; r++ {
				for _, g := range galaxies {
					if g.col == r {
						continue col_loop
					}
				}
				cols += expansion_factor - 1
			}
			fmt.Println("distance between", i+1, "and", j+1, ":", rows+cols, "(", rows, cols, ")")
			total_dist += rows + cols

		}
	}
	return total_dist
}

func main() {
	input := utils.NewStdinInput()
	fmt.Fprintln(os.Stderr, compute_galaxy_dists(input, 2))
}
