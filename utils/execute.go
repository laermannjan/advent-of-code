package utils

import (
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func GetFlags() (input string, only_one bool, only_two bool) {
	flag.StringVar(&input, "input", "./input.txt", "path to input")
	flag.BoolVar(&only_one, "one", false, "execute part one")
	flag.BoolVar(&only_two, "two", false, "execute part two")

	var example int
	flag.IntVar(&example, "example", 0, "example input")
	flag.Lookup("example").NoOptDefVal = "1"

	var quiet bool
	flag.BoolVarP(&quiet, "quiet", "q", false, "whether to show debug output")

	flag.Parse()

	if quiet {
		log.SetOutput(io.Discard)
	}

	if example != 0 {
		input = fmt.Sprintf("./example%d.txt", example)
		if _, err := os.Stat(input); example < 2 && errors.Is(err, os.ErrNotExist) {
			input = "./example.txt"
		}
	}

	return
}

func Execute(fn func(Input) interface{}, input string) (result interface{}, elapsed time.Duration) {

	start := time.Now()
	result = fn(&FileInput{FilePath: input})
	elapsed = time.Since(start)
	return
}

type Day struct {
	PartOne func(Input) interface{}
	PartTwo func(Input) interface{}
}

func (d Day) Run() {
	input, only_one, only_two := GetFlags()
	_, f, _, _ := runtime.Caller(1)
	input = filepath.Join(filepath.Dir(f), input)
	if !only_two {
		result, elapsed := Execute(d.PartOne, input)
		fmt.Printf("part one: %v (took: %v)\n", result, elapsed)
	}
	if !only_one {
		result, elapsed := Execute(d.PartTwo, input)
		fmt.Printf("part two: %v (took: %v)\n", result, elapsed)
	}

}
