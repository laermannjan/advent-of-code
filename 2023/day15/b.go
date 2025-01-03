package main

import (
	"fmt"
	"lj/utils"
	"os"
	"strings"
)

func hash(input string) int {
	cur := 0
	for _, ch := range input {
		cur += int(ch)
		cur *= 17
		cur %= 256
	}
	return cur
}

type Lens struct {
	label string
	focal int
}

type Box struct {
	lenses []Lens
}

func (b *Box) remove(lens Lens) {
	for i, l := range b.lenses {
		if l.label == lens.label {
			b.lenses = append(b.lenses[:i], b.lenses[i+1:]...)
			return
		}
	}
}

func (b *Box) insert(lens Lens) {
	for i, l := range b.lenses {
		if l.label == lens.label {
			b.lenses[i] = lens
			return
		}
	}
	b.lenses = append(b.lenses, lens)
}

func (b *Box) focusPower(boxId int) int {
	answer := 0
	for l, lens := range b.lenses {
		lens_power := (boxId + 1) * (l + 1) * lens.focal
		fmt.Println("lens", l, "power", lens_power)
		answer += lens_power
	}
	return answer
}

func main() {
	input := utils.NewStdinInput()
	instructions := strings.Split(input.LineSlice()[0], ",")
	boxes := [256]Box{}
	for _, inst_str := range instructions {
		if strings.Contains(inst_str, "-") {
			inst := []rune(inst_str)
			label := string(inst[:len(inst)-1])
			box := hash(label)
			boxes[box].remove(Lens{label: label})

			fmt.Println("after", inst_str, boxes[box])
		} else {
			inst := []rune(inst_str)
			label := string(inst[:len(inst)-2])
			box := hash(label)
			focal := utils.Atoi(string(inst[len(inst)-1]))
			boxes[box].insert(Lens{label: label, focal: focal})

			fmt.Println("after", inst_str, boxes[box])
		}
	}

	total := 0
	for b, box := range boxes {
		total += box.focusPower(b)
	}

	fmt.Fprintln(os.Stderr, total)
}
