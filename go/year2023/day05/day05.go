package day05

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"reflect"
	"slices"
	"strings"

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

type Range struct {
	start  int
	length int
}

func (r Range) String() string {
	return fmt.Sprintf("[%d<-%d->%d]", r.start, r.length, r.start+r.length-1)
}

func partA_old(input utils.Input) int {
	sections := input.SectionSlice()
	seeds := strings.Fields(strings.Split(sections[0], ":")[1])
	locations := []int{}

	src_dest_maps := []map[Range]int{}

	for _, section := range sections[1:] {
		src_dest_map := map[Range]int{}
		for _, mapping := range strings.Split(section, "\n")[1:] {

			fields := strings.Fields(mapping)
			dest := utils.Atoi(fields[0])
			source := utils.Atoi(fields[1])
			length := utils.Atoi(fields[2])

			src_dest_map[Range{source, source + length}] = dest - source
		}
		src_dest_maps = append(src_dest_maps, src_dest_map)
	}

	for _, seed := range seeds {
		n := utils.Atoi(seed)
		log.Println("seed:", n)
		for _, cur_map := range src_dest_maps {
			for cur_range, shift := range cur_map {
				if cur_range.start <= n && n <= cur_range.start+cur_range.length {
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
	sections := input.SectionSlice()
	seeds_str := strings.Fields(strings.Split(sections[0], ":")[1])
	seeds := []int{}
	for _, seed := range seeds_str {
		seeds = append(seeds, utils.Atoi(seed))
	}

	log.Printf("%-30s: %v", "initial", seeds)
	for _, section := range sections[1:] {
		mapped := map[int]bool{}

		name, _, _ := strings.Cut(section, ":")
		for _, m := range strings.Split(section, "\n")[1:] {

			fields := strings.Fields(m)
			dest := utils.Atoi(fields[0])
			source := utils.Atoi(fields[1])
			length := utils.Atoi(fields[2])
			log.Println(dest, source, length)

			for i, seed := range seeds {
				if !mapped[i] && source <= seed && seed < source+length {
					mapped[i] = true
					seeds[i] += dest - source
					log.Printf("seed=%v [%v, %v) -> %v", seed, source, source+length, seeds[i])
				}
			}

		}
		log.Printf("%-30s: %v", name, seeds)
	}
	min_loc := math.MaxInt
	for _, seed := range seeds {
		if seed < min_loc {
			min_loc = seed
		}
	}
	return min_loc
}

type Conversion struct {
	Source      int
	Destination int
	Length      int
}

func (c Conversion) String() string {
	source := Range{start: c.Source, length: c.Length}
	dest := Range{start: c.Destination, length: c.Length}
	return fmt.Sprintf("%v >> %v", source, dest)
}

type Map struct {
	From        string
	To          string
	Conversions []Conversion
}

func (m Map) String() string {
	parts := []string{fmt.Sprintf("%s -> %s", m.From, m.To)}
	for _, conv := range m.Conversions {
		parts = append(parts, conv.String())
	}

	return strings.Join(parts, "\n")
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

func parseSeedRanges(str string) (seeds []Range) {
	str, found := strings.CutPrefix(str, "seeds:")
	if !found {
		log.Fatal("could not parse seed ranges")
	}

	fields := strings.Fields(str)
	for i := 0; i < len(fields); i += 2 {
		start := utils.Atoi(fields[i])
		length := utils.Atoi(fields[i+1])
		seeds = append(seeds, Range{start: start, length: length})
	}
	slices.SortFunc(seeds, func(a, b Range) int { return cmp.Compare(a.start, b.start) })
	return
}

func parseMap(s string) (m Map) {
	lines := strings.Split(s, "\n")
	fields := strings.Split(strings.Split(lines[0], " ")[0], "-to-")
	m.From = fields[0]
	m.To = fields[1]
	for _, conv := range lines[1:] {
		fields := strings.Fields(conv)
		m.Conversions = append(m.Conversions, Conversion{
			Destination: utils.Atoi(fields[0]),
			Source:      utils.Atoi(fields[1]),
			Length:      utils.Atoi(fields[2]),
		})
	}
	log.Println("before", m.Conversions)
	slices.SortFunc(m.Conversions, func(a, b Conversion) int { return cmp.Compare(a.Source, b.Source) })
	log.Println("after", m.Conversions)

	return
}

func partB(input utils.Input) int {
	sections := input.SectionSlice()
	seeds := parseSeedRanges(sections[0])
	log.Println("seeds:", seeds)

	maps := []Map{}
	for _, section := range sections[1:] {
		maps = append(maps, parseMap(section))
	}

	for _, m := range maps {
		log.Println()
		log.Println("Map from", m.From, "to", m.To)
		log.Println("Starting seeds:", seeds)
		converted_seeds := []Range{}
		for _, seed := range seeds {
			log.Println("Examining new seed range:", seed)
			for _, conv := range m.Conversions {
				log.Println("checking conv rule", conv)
				if seed.start <= conv.Source+conv.Length && conv.Source < seed.start+seed.length {
					//overlap
					log.Printf("Overlap found between seed: %v conv: %v", seed, conv)
					offset := conv.Destination - conv.Source

					if seed.start < conv.Source {
						// pass through, as there is no conversion rule for this range
						passthrough_part := Range{start: seed.start, length: conv.Source - seed.start}
						log.Println("passthrough part:", passthrough_part)

						converted_seeds = append(converted_seeds, passthrough_part)
						log.Println("setting seed.start", seed.start, "to", conv.Source)
						seed.start = conv.Source
					}

					if seed.start+seed.length <= conv.Source+conv.Length {
						// remaining seed is completely within conversion range, convert entire seed
						converted_seed := Range{start: seed.start + offset, length: seed.length}
						log.Println("seed remainder entirely in conversion rule:", seed, "converted to", converted_seed)
						converted_seeds = append(converted_seeds, converted_seed)
						seed = Range{start: seed.start + seed.length, length: 0}
						break
					} else {
						// part of remaining seed is inside this convertion, but part outside and may be converted by a following converision
						log.Println("seed overflows conversion rule, adding covered range")

						// part inside this converison rule
						length_inside := conv.Source + conv.Length - seed.start
						converted_seed_part := Range{start: seed.start + offset, length: length_inside}
						converted_seeds = append(converted_seeds, converted_seed_part)
						log.Println("converting", Range{start: seed.start, length: length_inside}, "to", converted_seed_part)

						// remainder not covered by this conversion rule
						seed = Range{start: conv.Source + conv.Length, length: seed.start + seed.length - (conv.Source + conv.Length)}
						log.Println("reminaing seed", seed)
					}
				}
			}
			if seed.length > 0 {
				log.Println("seed", seed, "did not match any conversion rule, passing through")
				converted_seeds = append(converted_seeds, seed)
			}
		}

		log.Println("Converted seeds:", converted_seeds)
		slices.SortFunc(converted_seeds, func(a, b Range) int { return cmp.Compare(a.start, b.start) })
		log.Println("Sorted seeds:", converted_seeds)

		merged_seeds := []Range{}
		current := converted_seeds[0]
		for _, next := range converted_seeds[1:] {
			if current.start+current.length >= next.start {
				// overlapping or adjacent
				end := max(current.start+current.length, next.start+next.length)
				current.length = end - current.start
			} else {
				merged_seeds = append(merged_seeds, current)
				current = next
			}

		}
		merged_seeds = append(merged_seeds, current)
		log.Println("Merged seeds:", merged_seeds)

		seeds = merged_seeds
	}
	return seeds[0].start
}
