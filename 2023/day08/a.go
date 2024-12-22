package main

import (
	"fmt"
	"lj/utils"
	"os"
	"regexp"
)

type option struct {
	left  string
	right string
}

func main() {
	input := utils.NewStdinInput()

	lines := input.LineSlice()
	instructions := []rune(lines[0])

	network := map[string]option{}

	re := regexp.MustCompile(`(...) = \((...), (...)\)`)

	for _, line := range lines[2:] {
		match := re.FindStringSubmatch(line)
		network[match[1]] = option{left: match[2], right: match[3]}
	}

	current := "AAA"

	i := 0
	for current != "ZZZ" {
		inst := instructions[i%len(instructions)]

		switch inst {
		case 'L':
			current = network[current].left
		case 'R':
			current = network[current].right
		}
		fmt.Println("step:", i, "inst:", string(inst), "->", current)

		i++
	}
	fmt.Fprintln(os.Stderr, i)
}
