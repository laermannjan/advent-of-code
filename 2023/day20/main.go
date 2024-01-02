package main

import (
	"aoc-go/utils"
	"log"
	"slices"
	"strings"
)

type Pulse string

const (
	low  Pulse = "low"
	high Pulse = "high"
	null Pulse = "<null>"
)

type Transmission struct {
	from  string
	to    string
	pulse Pulse
}

type FlipFlop struct {
	state   Pulse
	outputs []string
}

func (m *FlipFlop) getOutputs() []string {
	return m.outputs
}

func (m *FlipFlop) output(t Transmission) []Transmission {
	pulse := null

	if t.pulse != high {
		switch m.state {
		case low:
			m.state = high
		case high:
			m.state = low
		}
		pulse = m.state
	}

	outs := []Transmission{}
	for _, out := range m.outputs {
		if pulse != null {
			outs = append(outs, Transmission{from: t.to, pulse: pulse, to: out})
		}
	}

	return outs
}

type Conjunction struct {
	states  map[string]Pulse
	outputs []string
}

func (m *Conjunction) getOutputs() []string {
	return m.outputs
}

func (m *Conjunction) output(t Transmission) []Transmission {
	m.states[t.from] = t.pulse

	// log.Println(m.states)

	pulse := low
	for _, state := range m.states {
		// log.Println("checking", state)
		if state == low {
			pulse = high
			// log.Println("setting pulse", pulse)
		}
	}
	// log.Println("pulse is", pulse)

	outs := []Transmission{}
	for _, out := range m.outputs {
		outs = append(outs, Transmission{from: t.to, pulse: pulse, to: out})
	}
	return outs
}

type Broadcaster struct {
	outputs []string
}

func (m *Broadcaster) getOutputs() []string {
	return m.outputs
}

func (m *Broadcaster) output(t Transmission) []Transmission {
	outs := []Transmission{}
	for _, out := range m.outputs {
		outs = append(outs, Transmission{from: t.to, pulse: t.pulse, to: out})
	}
	return outs
}

type Module interface {
	getOutputs() []string
	output(Transmission) []Transmission
}

func part1(input utils.Input) (answer interface{}) {
	modules := map[string]Module{}
	for line := range input.Lines() {
		con := strings.Split(line, " -> ")
		outputs := strings.Split(con[1], ", ")

		var module string
		if con[0] == "broadcaster" {
			module = con[0]
			modules[module] = &Broadcaster{outputs: outputs}

		} else {
			con_runes := []rune(con[0])
			module = string(con_runes[1:])
			switch con_runes[0] {
			case '%':
				modules[module] = &FlipFlop{state: low, outputs: outputs}
			case '&':
				modules[module] = &Conjunction{states: map[string]Pulse{}, outputs: outputs}
			}
		}
	}

	for k, v := range modules {
		for _, out := range v.getOutputs() {
			switch modules[out].(type) {
			case *Conjunction:
				modules[out].(*Conjunction).states[k] = low
			}
		}
	}

	for k, v := range modules {
		log.Printf("%v: %#v", k, v)
	}

	low_count := 0
	high_count := 0

	for i := 0; i < 1000; i++ {

		init := Transmission{
			from:  "button", // maybe this causes troubles with conjunctions
			to:    "broadcaster",
			pulse: low,
		}

		stack := []Transmission{init}

		log.Println("stack:", stack)
		for len(stack) > 0 {
			t := stack[0]

			switch t.pulse {
			case high:
				high_count++
			case low:
				low_count++
			}

			log.Printf("%v -%v-> %v", t.from, t.pulse, t.to)
			stack = stack[1:]

			if receiver, ok := modules[t.to]; ok {
				outputs := receiver.output(t)
				log.Println("adding outputs", outputs)
				stack = append(stack, outputs...)

			}
		}
	}

	log.Println("low", low_count, "high", high_count)

	return low_count * high_count
}

// could only solve this with some assumptions without basis
// the target "rx" has only one input, which is a Conjunction
// a conjuction sends a low pulse if all of its inputs have been sending a high pulse
// so I assume that its inputs send high pulses in intervals and we need to
// find the earliest button press where all these intervals coincide
func part2(input utils.Input) (answer interface{}) {
	modules := map[string]Module{}
	feed := ""
	cycles := map[string]int{}

	for line := range input.Lines() {
		con := strings.Split(line, " -> ")
		outputs := strings.Split(con[1], ", ")

		var module string
		if con[0] == "broadcaster" {
			module = con[0]
			modules[module] = &Broadcaster{outputs: outputs}

		} else {
			con_runes := []rune(con[0])
			module = string(con_runes[1:])
			switch con_runes[0] {
			case '%':
				modules[module] = &FlipFlop{state: low, outputs: outputs}
			case '&':
				modules[module] = &Conjunction{states: map[string]Pulse{}, outputs: outputs}
			}
		}
		if slices.Contains(outputs, "rx") {
			feed = module
		}

	}

	for k, v := range modules {
		for _, out := range v.getOutputs() {
			switch modules[out].(type) {
			case *Conjunction:
				modules[out].(*Conjunction).states[k] = low
			}
		}
	}

	for k, v := range modules {
		log.Printf("%v: %#v", k, v)
	}

	rx_count := 0

	i := 0
	for {
		i++
		// log.Println("i", i)

		init := Transmission{
			from:  "button", // maybe this causes troubles with conjunctions
			to:    "broadcaster",
			pulse: low,
		}

		stack := []Transmission{init}

		// log.Println("stack:", stack)
		for len(stack) > 0 {
			t := stack[0]

			if t.to == feed && t.pulse == high {
				cycles[t.from] = i - cycles[t.from]
			}

			// log.Printf("%v -%v-> %v", t.from, t.pulse, t.to)
			stack = stack[1:]

			if receiver, ok := modules[t.to]; ok {
				outputs := receiver.output(t)
				// log.Println("adding outputs", outputs)
				stack = append(stack, outputs...)

			}
		}
		if rx_count == 1 || i > 10000 {
			break
		}
	}

	log.Println("cycles", cycles, "rx", rx_count)

	cycle_lengths := []int{}
	for _, cycle := range cycles {
		cycle_lengths = append(cycle_lengths, cycle)
	}
	return utils.LCM(cycle_lengths...)
}

func main() {
	utils.Day{PartOne: part1, PartTwo: part2}.Run()
}
