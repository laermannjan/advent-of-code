package main

import (
	"aoc-go/utils"
	"log"
	"regexp"
	"strings"
)

type option struct {
	left  string
	right string
}

func part1(input utils.Input) interface{} {
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
		log.Println("step:", i, "inst:", string(inst), "->", current)

		i++
	}
	return i
}

// this doesn't work, it will take too long to simulate all the paths
func part2_old(input utils.Input) interface{} {
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

		log.Println("step:", i, "inst:", string(inst), "->", currents)

		i++
	}
	return i
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
func part2(input utils.Input) interface{} {
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

		log.Println("current:", current)

		for {
			inst := instructions[i%len(instructions)]
			switch inst {
			case 'L':
				current = network[current].left
			case 'R':
				current = network[current].right
			}
			log.Println("step:", i, "inst:", string(inst), "->", current)
			cycle_length++
			i++
			if strings.HasSuffix(current, "Z") {
				log.Println("found suffix in", current)
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
	log.Println("cycles:", cycles)
	// for c, next := range cycles[1:] {
	// 	prev := cycles[c]
	//
	// }
	lcm := utils.LCM(cycles...)

	return lcm
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
