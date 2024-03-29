package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Input interface {
	Lines() <-chan string
	LineSlice() []string
	RunesSlice() [][]rune
	Sections() <-chan string
	SectionSlice() []string
}

type FileInput struct {
	FilePath string
}

func (fi *FileInput) Lines() <-chan string {
	file, err := os.Open(fi.FilePath)
	if err != nil {
		panic(err)
	}
	// defer file.Close()  // TODO: why do I not need to close this, if I do nothing gets through the channel

	scanner := bufio.NewScanner(file)
	ch := make(chan string)

	go func() {
		defer close(ch)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}

func (fi *FileInput) LineSlice() []string {
	line_slice := []string{}
	for line := range fi.Lines() {
		line_slice = append(line_slice, line)
	}
	return line_slice
}

func (fi *FileInput) RunesSlice() [][]rune {
	runes_slice := [][]rune{}
	for line := range fi.Lines() {
		runes_slice = append(runes_slice, []rune(line))
	}
	return runes_slice
}

func (fi *FileInput) Sections() <-chan string {
	file, err := os.Open(fi.FilePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	ch := make(chan string)

	go func() {
		// defer file.Close()  // TODO: should I do this, or not?
		defer close(ch)
		var section string
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" { // match empty line
				ch <- strings.TrimSpace(section)
				section = ""
			} else {
				section += line + "\n"
			}
		}
		if len(section) > 0 {
			ch <- strings.TrimSpace(section)
		}
	}()

	return ch
}

func (fi *FileInput) SectionSlice() []string {
	para_slice := []string{}
	for para := range fi.Sections() {
		para_slice = append(para_slice, para)
	}
	return para_slice
}

func FromInputFile(year int, day int) Input {
	path := filepath.Join(os.Getenv("AOC_DATA_ROOT"), fmt.Sprintf("%d", year), "inputs", fmt.Sprintf("%02d", day)+".txt")
	return &FileInput{FilePath: path}
}

func FromExampleFile(year int, day int, part string) Input {
	path := filepath.Join(os.Getenv("AOC_DATA_ROOT"), fmt.Sprintf("%d", year), "examples", fmt.Sprintf("%02d-", day)+part+".txt")
	return &FileInput{FilePath: path}
}
