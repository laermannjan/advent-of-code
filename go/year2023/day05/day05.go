package day05

import (
	"fmt"
	"log"
	"math"
	"reflect"
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
	lower int
	upper int
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
				if cur_range.lower <= n && n <= cur_range.upper {
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

func partB(input utils.Input) int {
	return 0
}
