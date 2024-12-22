package main

import (
	"cmp"
	"fmt"
	"lj/utils"
	"math"
	"os"
	"slices"
	"strings"
)

func main_old() {
	input := utils.NewStdinInput()

	sections := input.SectionSlice()
	seeds := strings.Fields(strings.Split(sections[0], ":")[1])
	locations := []int{}

	src_dest_maps := []map[Interval]int{}

	for _, section := range sections[1:] {
		src_dest_map := map[Interval]int{}
		for _, mapping := range strings.Split(section, "\n")[1:] {

			fields := strings.Fields(mapping)
			dest := utils.Atoi(fields[0])
			source := utils.Atoi(fields[1])
			length := utils.Atoi(fields[2])

			src_dest_map[Interval{source, source + length}] = dest - source
		}
		src_dest_maps = append(src_dest_maps, src_dest_map)
	}

	for _, seed := range seeds {
		n := utils.Atoi(seed)
		fmt.Println("seed:", n)
		for _, cur_map := range src_dest_maps {
			for cur_range, shift := range cur_map {
				if cur_range.start <= n && n <= cur_range.end {
					n += shift
					break
				}
			}

			fmt.Println("mapped to:", n)
		}
		locations = append(locations, n)

	}

	min_loc := math.MaxInt
	for _, loc := range locations {
		if loc < min_loc {
			min_loc = loc
		}
	}

	fmt.Fprintln(os.Stderr, min_loc)
}

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
	seeds := parseSeeds(sections[0])
	fmt.Println("Initial seeds:", seeds)

	maps := []Map{}
	for _, section := range sections[1:] {
		maps = append(maps, parseMap(section))
	}

	for _, m := range maps {
		fmt.Println("\n\nMap:", m)
		for i, seed := range seeds {
			fmt.Println("\nseed:", seed)
			for _, rule := range m.rules {
				fmt.Print("rule:", rule)
				if rule.start <= seed && seed < rule.end {
					seeds[i] += rule.offset
					fmt.Println("converting:", seed, "->", seeds[i])
					break
				} else {
					fmt.Println("no overlap")
				}
			}
		}

	}
	fmt.Fprintln(os.Stderr, slices.Min(seeds))
}
