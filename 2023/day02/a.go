package main

import (
	"fmt"
	"lj/utils"
	"os"
	"strings"
)

func main() {
	input := utils.NewStdinInput()

	max_cubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
game_loop:
	for line := range input.Lines() {
		parts := strings.Split(line, ":")
		game_id := utils.Atoi(strings.TrimPrefix(parts[0], "Game "))

		game := strings.Split(parts[1], ";")

		for _, set := range game {
			cubes := strings.Split(set, ",")
			for _, color := range cubes {
				v := strings.Fields(color)
				count := utils.Atoi(v[0])
				fmt.Println(v[1], count, max_cubes[v[1]])
				if count > max_cubes[v[1]] {
					continue game_loop
				}
			}
		}
		sum += game_id
	}
	fmt.Fprintln(os.Stderr, sum)
}
