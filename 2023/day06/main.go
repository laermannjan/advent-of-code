package main

import (
	"aoc-go/utils"
	"log"
	"math"
	"strings"
)

type Race struct {
	time   int
	record int
}

func parseRaces(input []string) (races []Race) {
	_times, _ := strings.CutPrefix(input[0], "Time:")
	times := utils.ParseInts(_times)
	_distances, _ := strings.CutPrefix(input[1], "Distance:")
	distances := utils.ParseInts(_distances)

	for i := 0; i < len(times); i++ {
		races = append(races, Race{time: times[i], record: distances[i]})
	}

	return
}

// Assume t := race time, d := race record distance, v := boat speed, s := charge time, d'(s) := boat's distance with given charge time
// Boat increases future speed by 1 for each millisecond we wait => v = s
// d'(s) = v * (t - s)
// d'(s) = s * (t - s)
// d'(s) = ts - s^2
// We want d'(s) > d, i.e. go farther than the record
// We solve for the zero points, where d = ts - s^2
// 0 = ts - s^2 - d
// 0 = s^2 - ts + d  // pq-formula to solve quadratic eq. (x^2 + px + q = 0 => -(p/2) +- sqrt([p/2]^2 - q))
// zero points are lower and upper bounds (non-inclusive) for s, in which we beat the race record
func part1(input utils.Input) interface{} {
	races := parseRaces(input.LineSlice())

	possible_wins := 0
	for _, r := range races {
		t := float64(r.time)
		d := float64(r.record)

		// zero points
		zp1 := t/2 - math.Sqrt(math.Pow(t/2, 2)-float64(d))
		zp2 := t/2 + math.Sqrt(math.Pow(t/2, 2)-float64(d))

		min_s := int(math.Ceil(zp1 + 1e-12))  // adding some margin to round up if we exactly reach d
		max_s := int(math.Floor(zp2 - 1e-12)) // subtracting some margin to round down ...

		log.Printf("{%.2f %.2f} -> {%d %d}", zp1, zp2, min_s, max_s)
		wins := max_s - min_s + 1
		log.Println("race:", r, "possible wins:", wins)

		if possible_wins == 0 {
			possible_wins = wins
		} else {
			possible_wins *= wins
		}
	}

	return possible_wins
}

func part2(input utils.Input) interface{} {
	lines := input.LineSlice()
	t_str, _ := strings.CutPrefix(lines[0], "Time:")
	t := float64(utils.Atoi(strings.Join(strings.Fields(t_str), "")))

	d_str, _ := strings.CutPrefix(lines[1], "Distance:")
	d := float64(utils.Atoi(strings.Join(strings.Fields(d_str), "")))

	zp1 := t/2 - math.Sqrt(math.Pow(t/2, 2)-float64(d))
	zp2 := t/2 + math.Sqrt(math.Pow(t/2, 2)-float64(d))

	min_s := int(math.Ceil(zp1 + 1e-12))  // adding some margin to round up if we exactly reach d
	max_s := int(math.Floor(zp2 - 1e-12)) // subtracting some margin to round down ...

	log.Printf("{%.2f %.2f} -> {%d %d}", zp1, zp2, min_s, max_s)
	wins := max_s - min_s + 1

	return wins

}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
