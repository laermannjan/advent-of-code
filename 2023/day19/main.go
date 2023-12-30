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

type bounds struct {
	lower int
	upper int
}

type JointCondition struct {
	x bounds
	m bounds
	a bounds
	s bounds
}

func part2(input utils.Input) (answer interface{}) {
	inp := input.SectionSlice()
	workflows_ := inp[0]

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

	chains := [][]string{{"in"}}
	combos := []JointCondition{
		{
			x: bounds{lower: 1, upper: 4000},
			m: bounds{lower: 1, upper: 4000},
			a: bounds{lower: 1, upper: 4000},
			s: bounds{lower: 1, upper: 4000},
		},
	}

	i := 0
	for i < len(chains) {

		this_chain := chains[i]
		log.Println("this", this_chain, "\tchains", chains)
		wf := this_chain[len(chains[i])-1]
		if wf == "A" {
			i++
			continue
		}

		new_chains := [][]string{}
		new_combos := []JointCondition{}

		next_combo := combos[i]

		for _, cond := range workflows[wf] {
			this_combo := next_combo
			log.Println("looking at", wf, cond, "new chains", new_chains)
			// log.Println("new chains before append", new_chains)

			new_chain := make([]string, len(this_chain))
			copy(new_chain, this_chain)

			switch cond.thing {
			case "x":
				switch cond.compare {
				case "<":
					this_combo.x.upper = min(this_combo.x.upper, cond.value-1)
					next_combo.x.lower = max(this_combo.x.lower, cond.value)
				case ">":
					this_combo.x.lower = max(this_combo.x.lower, cond.value+1)
					next_combo.x.upper = min(this_combo.x.upper, cond.value)
				}

			case "m":
				switch cond.compare {
				case "<":
					this_combo.m.upper = min(this_combo.m.upper, cond.value-1)
					next_combo.m.lower = max(this_combo.m.lower, cond.value)
				case ">":
					this_combo.m.lower = max(this_combo.m.lower, cond.value+1)
					next_combo.m.upper = min(this_combo.m.upper, cond.value)
				}
			case "a":
				switch cond.compare {
				case "<":
					this_combo.a.upper = min(this_combo.a.upper, cond.value-1)
					next_combo.a.lower = max(this_combo.a.lower, cond.value)
				case ">":
					this_combo.a.lower = max(this_combo.a.lower, cond.value+1)
					next_combo.a.upper = min(this_combo.a.upper, cond.value)
				}
			case "s":
				switch cond.compare {
				case "<":
					this_combo.s.upper = min(this_combo.s.upper, cond.value-1)
					next_combo.s.lower = max(this_combo.s.lower, cond.value)
				case ">":
					this_combo.s.lower = max(this_combo.s.lower, cond.value+1)
					next_combo.s.upper = min(this_combo.s.upper, cond.value)
				}
			}
			if cond.target != "R" {
				log.Println("adding combo", this_combo)
				new_combos = append(new_combos, this_combo)
				new_chain = append(new_chain, cond.target)
				log.Println("new chain with target", new_chain, new_chains)
				new_chains = append(new_chains, new_chain)
				log.Println("new_chains", new_chains)
			}
		}
		log.Println(chains[:i], "\t", new_chains, "\t", chains[i+1:])
		log.Println(combos[:i], "\t", new_combos, "\t", combos[i+1:])
		// log.Println("len new chains", len(new_chains))
		chains = append(append(append([][]string{}, chains[:i]...), new_chains...), chains[i+1:]...)
		combos = append(append(append([]JointCondition{}, combos[:i]...), new_combos...), combos[i+1:]...)

	}

	log.Println()
	for i := 0; i < len(chains); i++ {
		xr := combos[i].x.upper - combos[i].x.lower + 1
		mr := combos[i].m.upper - combos[i].m.lower + 1
		ar := combos[i].a.upper - combos[i].a.lower + 1
		sr := combos[i].s.upper - combos[i].s.lower + 1

		log.Println(chains[i], combos[i], xr*mr*ar*sr)
	}

	total := 0
	for _, combo := range combos {
		xr := combo.x.upper - combo.x.lower + 1
		mr := combo.m.upper - combo.m.lower + 1
		ar := combo.a.upper - combo.a.lower + 1
		sr := combo.s.upper - combo.s.lower + 1

		total += xr * mr * ar * sr
	}

	return total
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
