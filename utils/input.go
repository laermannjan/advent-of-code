package utils

import (
	"bufio"
	"os"
	"strings"
)

type Input interface {
	Lines() <-chan string
	LineSlice() []string
	RunesSlice() [][]rune
	Sections() <-chan string
	SectionSlice() []string
}

type StdinInput struct {
	lines []string
}

func NewStdinInput() *StdinInput {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &StdinInput{lines: lines}
}

func (si *StdinInput) Lines() <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, line := range si.lines {
			ch <- line
		}
	}()
	return ch
}

func (si *StdinInput) LineSlice() []string {
	return si.lines
}

func (si *StdinInput) RunesSlice() [][]rune {
	var runes [][]rune
	for _, line := range si.lines {
		runes = append(runes, []rune(line))
	}
	return runes
}

func (si *StdinInput) Sections() <-chan string {
	return sectionsFromLines(si.Lines())
}

func (si *StdinInput) SectionSlice() []string {
	return collectSections(si.Sections())
}

func collectSections(ch <-chan string) []string {
	var sections []string
	for section := range ch {
		sections = append(sections, section)
	}
	return sections
}

func sectionsFromLines(lines <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		var section strings.Builder
		for line := range lines {
			if line == "" {
				if section.Len() > 0 {
					ch <- strings.TrimSpace(section.String())
					section.Reset()
				}
			} else {
				section.WriteString(line + "\n")
			}
		}
		if section.Len() > 0 {
			ch <- strings.TrimSpace(section.String())
		}
	}()
	return ch
}
