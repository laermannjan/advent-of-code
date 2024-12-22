package main

import (
	"fmt"
	"lj/utils"
	"os"
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

	// fmt.Println(m.states)

	pulse := low
	for _, state := range m.states {
		// fmt.Println("checking", state)
		if state == low {
			pulse = high
			// fmt.Println("setting pulse", pulse)
		}
	}
	// fmt.Println("pulse is", pulse)

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

func main() {
	input := utils.NewStdinInput()
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
		fmt.Printf("%v: %#v", k, v)
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

		fmt.Println("stack:", stack)
		for len(stack) > 0 {
			t := stack[0]

			switch t.pulse {
			case high:
				high_count++
			case low:
				low_count++
			}

			fmt.Printf("%v -%v-> %v", t.from, t.pulse, t.to)
			stack = stack[1:]

			if receiver, ok := modules[t.to]; ok {
				outputs := receiver.output(t)
				fmt.Println("adding outputs", outputs)
				stack = append(stack, outputs...)

			}
		}
	}

	fmt.Println("low", low_count, "high", high_count)

	fmt.Fprintln(os.Stderr, low_count*high_count)
}
