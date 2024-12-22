package main

import (
	"fmt"
	"lj/utils"
	"os"
	"regexp"
	"strings"
)

type Condition struct {
	thing   string
	compare string
	value   int
	target  string
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

func main() {
	input := utils.NewStdinInput()
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
			fmt.Println("condition", condition)
			match := workflow_re.FindStringSubmatch(condition)
			if len(match) > 0 {
				workflows[components[0]] = append(workflows[components[0]], Condition{thing: match[1], compare: match[2], value: utils.Atoi(match[3]), target: match[4]})
			} else {
				workflows[components[0]] = append(workflows[components[0]], Condition{thing: "", compare: "", value: 0, target: condition})
			}
			fmt.Printf("%+v", workflows[components[0]])

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
		fmt.Println("this", this_chain, "\tchains", chains)
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
			fmt.Println("looking at", wf, cond, "new chains", new_chains)
			// fmt.Println("new chains before append", new_chains)

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
				fmt.Println("adding combo", this_combo)
				new_combos = append(new_combos, this_combo)
				new_chain = append(new_chain, cond.target)
				fmt.Println("new chain with target", new_chain, new_chains)
				new_chains = append(new_chains, new_chain)
				fmt.Println("new_chains", new_chains)
			}
		}
		fmt.Println(chains[:i], "\t", new_chains, "\t", chains[i+1:])
		fmt.Println(combos[:i], "\t", new_combos, "\t", combos[i+1:])
		// fmt.Println("len new chains", len(new_chains))
		chains = append(append(append([][]string{}, chains[:i]...), new_chains...), chains[i+1:]...)
		combos = append(append(append([]JointCondition{}, combos[:i]...), new_combos...), combos[i+1:]...)

	}

	fmt.Println()
	for i := 0; i < len(chains); i++ {
		xr := combos[i].x.upper - combos[i].x.lower + 1
		mr := combos[i].m.upper - combos[i].m.lower + 1
		ar := combos[i].a.upper - combos[i].a.lower + 1
		sr := combos[i].s.upper - combos[i].s.lower + 1

		fmt.Println(chains[i], combos[i], xr*mr*ar*sr)
	}

	total := 0
	for _, combo := range combos {
		xr := combo.x.upper - combo.x.lower + 1
		mr := combo.m.upper - combo.m.lower + 1
		ar := combo.a.upper - combo.a.lower + 1
		sr := combo.s.upper - combo.s.lower + 1

		total += xr * mr * ar * sr
	}

	fmt.Fprintln(os.Stderr, total)
}
