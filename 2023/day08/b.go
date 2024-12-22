package main

import (
	"fmt"
	"lj/utils"
	"os"
	"regexp"
	"strings"
)

type option struct {
	left  string
	right string
}

// this doesn't work, it will take too long to simulate all the paths
func main_old() {
	input := utils.NewStdinInput()

	lines := input.LineSlice()
	instructions := []rune(lines[0])

	network := map[string]option{}

	re := regexp.MustCompile(`(...) = \((...), (...)\)`)

	currents := []string{}
	for _, line := range lines[2:] {
		match := re.FindStringSubmatch(line)
		network[match[1]] = option{left: match[2], right: match[3]}
		if strings.HasSuffix(match[1], "A") {
			currents = append(currents, match[1])
		}
	}

	i := 0
	done := false
	for !done {
		inst := instructions[i%len(instructions)]

		done = true
		for c, current := range currents {
			switch inst {
			case 'L':
				currents[c] = network[current].left
			case 'R':
				currents[c] = network[current].right
			}
			done = done && strings.HasSuffix(currents[c], "Z")
		}

		fmt.Println("step:", i, "inst:", string(inst), "->", currents)

		i++
	}
	fmt.Fprintln(os.Stderr, i)
}

// Each starting node (..A) maps to a single destination node (..Z)
// Each path *seems* to be a cycle, the number of steps it takes from
// start -> first (..Z) is the same as all the following (..Z) -> (..Z)
// each individual path therefore only ever sees the same (..Z) node
// That means, the paths will visit their destination at stable frequencies
// Finding the least common multiple of all the path frequencies will be
// the step when they all meet up.
// the stable frequency was simply an assumption after observing this
// on the example input and printing out the first few cycles on the input
// it worked out, but couldn't get this from the problem description
func main() {
	input := utils.NewStdinInput()

	lines := input.LineSlice()
	instructions := []rune(lines[0])

	network := map[string]option{}

	re := regexp.MustCompile(`(...) = \((...), (...)\)`)

	currents := []string{}
	for _, line := range lines[2:] {
		match := re.FindStringSubmatch(line)
		network[match[1]] = option{left: match[2], right: match[3]}
		if strings.HasSuffix(match[1], "A") {
			currents = append(currents, match[1])
		}
	}

	cycles := []int{}
	for _, current := range currents {
		first_z := ""
		i := 0
		var cycle_length int

		fmt.Println("current:", current)

		for {
			inst := instructions[i%len(instructions)]
			switch inst {
			case 'L':
				current = network[current].left
			case 'R':
				current = network[current].right
			}
			fmt.Println("step:", i, "inst:", string(inst), "->", current)
			cycle_length++
			i++
			if strings.HasSuffix(current, "Z") {
				fmt.Println("found suffix in", current)
				if first_z == "" {
					first_z = current
					cycle_length = 0
				} else if current == first_z {
					break
				}

			}
		}
		cycles = append(cycles, cycle_length)
	}
	fmt.Println("cycles:", cycles)
	// for c, next := range cycles[1:] {
	// 	prev := cycles[c]
	//
	// }
	lcm := utils.LCM(cycles...)

	fmt.Fprintln(os.Stderr, lcm)
}
