package day05

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/laermannjan/advent-of-code/go/utils"
	"github.com/spf13/cobra"
)

func ACmd() *cobra.Command {
	part := "a"
	return &cobra.Command{
		Use:   part,
		Short: "Part " + part,
		Run: func(cmd *cobra.Command, _ []string) {
			type PkgMark struct{}
			year, day := utils.GetYearDay(reflect.TypeOf(PkgMark{}).PkgPath())

			solveExample, err := cmd.Flags().GetBool("example")
			if err != nil {
				log.Fatal(err)
			}

			var input utils.Input
			if solveExample {
				input = utils.FromExampleFile(year, day, part)
			} else {
				input = utils.FromInputFile(year, day)
			}

			fmt.Printf("Answer: %d\n", partA(input))
		},
	}
}

func BCmd() *cobra.Command {
	part := "b"
	return &cobra.Command{
		Use:   part,
		Short: "Part " + part,
		Run: func(cmd *cobra.Command, _ []string) {
			type PkgMark struct{}
			year, day := utils.GetYearDay(reflect.TypeOf(PkgMark{}).PkgPath())

			solveExample, err := cmd.Flags().GetBool("example")
			if err != nil {
				log.Fatal(err)
			}

			var input utils.Input
			if solveExample {
				input = utils.FromExampleFile(year, day, part)
			} else {
				input = utils.FromInputFile(year, day)
			}

			fmt.Printf("Answer: %d\n", partB(input))
		},
	}
}

func partA_old(input utils.Input) int {
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
		log.Println("seed:", n)
		for _, cur_map := range src_dest_maps {
			for cur_range, shift := range cur_map {
				if cur_range.start <= n && n <= cur_range.end {
					n += shift
					break
				}
			}

			log.Println("mapped to:", n)
		}
		locations = append(locations, n)

	}

	min_loc := math.MaxInt
	for _, loc := range locations {
		if loc < min_loc {
			min_loc = loc
		}

	}

	return min_loc
}

func partA(input utils.Input) int {
	start := time.Now()
	sections := input.SectionSlice()
	seeds := parseSeeds(sections[0])
	log.Println("Initial seeds:", seeds)

	maps := []Map{}
	for _, section := range sections[1:] {
		maps = append(maps, parseMap(section))
	}

	for _, m := range maps {
		log.Println("\n\nMap:", m)
		for i, seed := range seeds {
			log.Println("\nseed:", seed)
			for _, rule := range m.rules {
				log.Print("rule:", rule)
				if rule.start <= seed && seed < rule.end {
					seeds[i] += rule.offset
					log.Println("converting:", seed, "->", seeds[i])
					break
				} else {
					log.Println("no overlap")
				}
			}
		}

	}
	fmt.Println("took:", time.Since(start))

	return slices.Min(seeds)
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
		log.Fatal("could not parse seeds")
	}
	for _, s := range strings.Fields(str) {
		seeds = append(seeds, utils.Atoi(s))
	}
	return
}

func parseSeedRanges(str string) (seeds []Interval) {
	str, found := strings.CutPrefix(str, "seeds:")
	if !found {
		log.Fatal("could not parse seed ranges")
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

func partB(input utils.Input) int {
	start := time.Now()

	sections := input.SectionSlice()
	seeds := parseSeedRanges(sections[0])
	log.Println("Initial seeds:", seeds)

	maps := []Map{}
	for _, section := range sections[1:] {
		maps = append(maps, parseMap(section))
	}

	for _, m := range maps {
		log.Println("\n\nMap:", m)
		log.Println("seeds:", seeds)
		converted_seeds := []Interval{}
		for _, seed := range seeds {
			log.Println()
			log.Println("seed:", seed)
			for _, rule := range m.rules {
				log.Println("\nrule", rule)

				if seed.start > rule.end || rule.start > seed.end {
					log.Println("no overlap")
					continue
				}

				if seed.start < rule.start {
					passthrough_seed_part := Interval{start: seed.start, end: rule.start}
					log.Println("seed.start < rule.start; passthrough:", passthrough_seed_part)

					converted_seeds = append(converted_seeds, passthrough_seed_part)
					seed.start = rule.start
					log.Println("remaining seed:", seed)
				}

				if seed.end <= rule.end {
					//seed fits into rule interval
					converted_seed := seed.convert(rule.offset)
					log.Println("converting (->seed.end):", seed, "->", converted_seed)
					converted_seeds = append(converted_seeds, converted_seed)
					seed.start = seed.end
					break
				} else {
					//seed overflows rule interval
					seed_part := Interval{start: seed.start, end: rule.end}
					converted_seed_part := seed_part.convert(rule.offset)
					converted_seeds = append(converted_seeds, converted_seed_part)
					log.Println("converting (->rule.end):", seed_part, "->", converted_seed_part)

					seed = Interval{start: rule.end, end: seed.end}
					log.Println("reminaing seed:", seed)
				}
			}
			if seed.start < seed.end {
				log.Println("no rule matched; passthrough:", seed)
				converted_seeds = append(converted_seeds, seed)
			}
		}

		log.Println("\nconverted seeds:", converted_seeds)
		slices.SortFunc(converted_seeds, func(a, b Interval) int { return cmp.Compare(a.start, b.start) })
		log.Println("sorted seeds:", converted_seeds)

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
		log.Println("merged seeds:", merged_seeds)

		seeds = merged_seeds
	}
	fmt.Println("took:", time.Since(start))
	return seeds[0].start
}
