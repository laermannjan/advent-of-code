package main

import (
	"aoc-go/utils"
	"cmp"
	"log"
	"regexp"
	"strings"
)

type Condition struct {
	thing   string
	compare string
	value   int
	target  string
}

func part1(input utils.Input) (answer interface{}) {
	inp := input.SectionSlice()
	workflows_ := inp[0]
	parts := inp[1]

	workflows := map[string][]Condition{
		"A": {},
		"R": {},
	}

	workflow_re := regexp.MustCompile(`([a-zA-Z]+)(\D+)(\d+):([a-zA-Z]+)`)
	for _, line := range strings.Split(workflows_, "\n") {
		components := strings.Split(line, "{")
		components[1] = strings.TrimSuffix(components[1], "}")
		for _, condition := range strings.Split(components[1], ",") {
			log.Println("condition", condition)
			match := workflow_re.FindStringSubmatch(condition)
			if len(match) > 0 {
				workflows[components[0]] = append(workflows[components[0]], Condition{thing: match[1], compare: match[2], value: utils.Atoi(match[3]), target: match[4]})
			} else {
				workflows[components[0]] = append(workflows[components[0]], Condition{thing: "", compare: "", value: 0, target: condition})
			}
			log.Printf("%+v", workflows[components[0]])

		}
	}

	parts_re := regexp.MustCompile(`x=(\d+),m=(\d+),a=(\d+),s=(\d+)`)

	sum := 0
	for _, line := range strings.Split(parts, "\n") {
		match := parts_re.FindStringSubmatch(line)
		vals := map[string]int{
			"x": utils.Atoi(match[1]),
			"m": utils.Atoi(match[2]),
			"a": utils.Atoi(match[3]),
			"s": utils.Atoi(match[4]),
		}

		log.Printf("%+v", vals)
		for wf := "in"; wf != "A" && wf != "R"; {
			conditions := workflows[wf]
			log.Println("workflow", wf, conditions)

		condition_loop:
			for _, condition := range conditions {

				log.Println("condition", condition, "part value", vals[condition.thing])

				comp := cmp.Compare(vals[condition.thing], condition.value)
				if condition.compare == "" || (condition.compare == "<" && comp < 0) || (condition.compare == ">" && comp > 0) {
					log.Println("move workflow", wf, "->", condition.target)
					wf = condition.target
					break condition_loop
				}

			}
			if wf == "A" {
				sum += vals["x"] + vals["m"] + vals["a"] + vals["s"]
			}
		}
	}

	return sum
}

func part2(input utils.Input) (answer interface{}) {
	return
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
