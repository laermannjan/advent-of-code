package main

import (
	"fmt"
	"lj/utils"
	"os"
	"strings"
)

func main() {
	input := utils.NewStdinInput()

	sum := 0
	for line := range input.Lines() {
		needed_cubes := map[string]int{}
		parts := strings.Split(line, ":")
		game := strings.Split(parts[1], ";")

		for _, set := range game {
			cubes := strings.Split(set, ",")
			for _, color := range cubes {
				v := strings.Fields(color)
				count := utils.Atoi(v[0])
				fmt.Println(v[1], count, needed_cubes[v[1]])

				needed_cubes[v[1]] = max(needed_cubes[v[1]], count)
			}
		}

		product := 1
		for _, v := range needed_cubes {
			product *= v
		}

		sum += product
	}
	fmt.Fprintln(os.Stderr, sum)
}
