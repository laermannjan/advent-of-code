package utils

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"log"
	"path/filepath"
	"runtime"
	"time"
)

func GetFlags() (input string, part string) {
	flag.StringVar(&input, "input", "", "input file path")
	flag.StringVar(&part, "part", "one", "run part [one|two]")

	var verbose bool
	flag.BoolVarP(&verbose, "verbose", "v", false, "whether to show debug output")

	flag.Parse()

	if !verbose {
		log.SetOutput(io.Discard)
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
	input, part := GetFlags()
	_, f, _, _ := runtime.Caller(1)

	if !filepath.IsAbs(input) {
		input = filepath.Join(filepath.Dir(f), input)
	}
	var fn func(Input) interface{}
	if part == "one" {
		fn = d.PartOne
	} else if part == "two" {
		fn = d.PartTwo
	}
	result, elapsed := Execute(fn, input)
	fmt.Printf("part %s: %v (took: %v)\n", part, result, elapsed)

}
