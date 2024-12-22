package main

import (
	"cmp"
	"fmt"
	"lj/utils"
	"os"
	"slices"
	"strings"
)

type Interval struct {
	start int
	end   int
}

func (i Interval) convert(offset int) Interval {
	return Interval{start: i.start + offset, end: i.end + offset}
}

func (r Interval) String() string {
	return fmt.Sprintf("[%d,%d)", r.start, r.end)
}

type Rule struct {
	start  int
	end    int
	offset int
}

func (c Rule) String() string {
	sign := ""
	if c.offset > 0 {
		sign = "+"
	}
	return fmt.Sprintf("[%d,%d) -> %s%v", c.start, c.end, sign, c.offset)
}

type Map struct {
	from  string
	to    string
	rules []Rule
}

func (m Map) String() string {
	return fmt.Sprintf("%s -> %s", m.from, m.to)
}

func parseSeeds(str string) (seeds []int) {
	str, found := strings.CutPrefix(str, "seeds:")
	if !found {
		fmt.Fprintln(os.Stderr, "could not parse seeds")
		os.Exit(1)
	}
	for _, s := range strings.Fields(str) {
		seeds = append(seeds, utils.Atoi(s))
	}
	return
}

func parseSeedRanges(str string) (seeds []Interval) {
	str, found := strings.CutPrefix(str, "seeds:")
	if !found {
		fmt.Fprintln(os.Stderr, "could not parse seed ranges")
		os.Exit(1)
	}

	fields := strings.Fields(str)
	for i := 0; i < len(fields); i += 2 {
		start := utils.Atoi(fields[i])
		length := utils.Atoi(fields[i+1])
		seeds = append(seeds, Interval{start: start, end: start + length})
	}
	slices.SortFunc(seeds, func(a, b Interval) int { return cmp.Compare(a.start, b.start) })
	return
}

func parseMap(s string) (m Map) {
	lines := strings.Split(s, "\n")
	fields := strings.Split(strings.Split(lines[0], " ")[0], "-to-")
	m.from = fields[0]
	m.to = fields[1]
	for _, conv := range lines[1:] {
		fields := strings.Fields(conv)

		dest := utils.Atoi(fields[0])
		source := utils.Atoi(fields[1])
		length := utils.Atoi(fields[2])

		m.rules = append(m.rules, Rule{
			start:  source,
			end:    source + length,
			offset: dest - source,
		})
	}
	slices.SortFunc(m.rules, func(a, b Rule) int { return cmp.Compare(a.start, b.start) })
	return
}

func main() {
	input := utils.NewStdinInput()

	sections := input.SectionSlice()
	seeds := parseSeedRanges(sections[0])
	fmt.Println("Initial seeds:", seeds)

	maps := []Map{}
	for _, section := range sections[1:] {
		maps = append(maps, parseMap(section))
	}

	for _, m := range maps {
		fmt.Println("\n\nMap:", m)
		fmt.Println("seeds:", seeds)
		converted_seeds := []Interval{}
		for _, seed := range seeds {
			fmt.Println()
			fmt.Println("seed:", seed)
			for _, rule := range m.rules {
				fmt.Println("\nrule", rule)

				if seed.start > rule.end || rule.start > seed.end {
					fmt.Println("no overlap")
					continue
				}

				if seed.start < rule.start {
					passthrough_seed_part := Interval{start: seed.start, end: rule.start}
					fmt.Println("seed.start < rule.start; passthrough:", passthrough_seed_part)

					converted_seeds = append(converted_seeds, passthrough_seed_part)
					seed.start = rule.start
					fmt.Println("remaining seed:", seed)
				}

				if seed.end <= rule.end {
					// seed fits into rule interval
					converted_seed := seed.convert(rule.offset)
					fmt.Println("converting (->seed.end):", seed, "->", converted_seed)
					converted_seeds = append(converted_seeds, converted_seed)
					seed.start = seed.end
					break
				} else {
					// seed overflows rule interval
					seed_part := Interval{start: seed.start, end: rule.end}
					converted_seed_part := seed_part.convert(rule.offset)
					converted_seeds = append(converted_seeds, converted_seed_part)
					fmt.Println("converting (->rule.end):", seed_part, "->", converted_seed_part)

					seed = Interval{start: rule.end, end: seed.end}
					fmt.Println("reminaing seed:", seed)
				}
			}
			if seed.start < seed.end {
				fmt.Println("no rule matched; passthrough:", seed)
				converted_seeds = append(converted_seeds, seed)
			}
		}

		fmt.Println("\nconverted seeds:", converted_seeds)
		slices.SortFunc(converted_seeds, func(a, b Interval) int { return cmp.Compare(a.start, b.start) })
		fmt.Println("sorted seeds:", converted_seeds)

		merged_seeds := []Interval{}
		current := converted_seeds[0]
		for _, next := range converted_seeds[1:] {
			if current.end >= next.start {
				// overlapping or adjacent
				current.end = max(current.end, next.end)
			} else {
				merged_seeds = append(merged_seeds, current)
				current = next
			}
		}
		merged_seeds = append(merged_seeds, current)
		fmt.Println("merged seeds:", merged_seeds)

		seeds = merged_seeds
	}
	fmt.Fprintln(os.Stderr, seeds[0].start)
}
