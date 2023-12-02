package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Input interface {
	Lines() <-chan string
	Paragraphs() <-chan string
}

type FileInput struct {
	filePath string
}

func (fi *FileInput) Lines() <-chan string {
	file, err := os.Open(fi.filePath)
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

func (fi *FileInput) Paragraphs() <-chan string {
	file, err := os.Open(fi.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ch := make(chan string)

	go func() {
		defer close(ch)
		var paragraph string
		for scanner.Scan() {
			if scanner.Text() == "" {
				ch <- paragraph
				paragraph = ""
			} else {
				paragraph += scanner.Text() + "\n"
			}
		}
	}()

	return ch
}

func FromInputFile(year int, day int) Input {
	path := filepath.Join(os.Getenv("AOC_DATA_ROOT"), fmt.Sprintf("%d", year), "inputs", fmt.Sprintf("%02d", day)+".txt")
	return &FileInput{filePath: path}
}

func FromExampleFile(year int, day int, part string) Input {
	path := filepath.Join(os.Getenv("AOC_DATA_ROOT"), fmt.Sprintf("%d", year), "examples", fmt.Sprintf("%02d-", day)+part+".txt")
	return &FileInput{filePath: path}
}
