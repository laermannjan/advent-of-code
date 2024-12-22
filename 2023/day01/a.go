package main

import (
	"fmt"
	"lj/utils"
	"os"
	"strconv"
	"unicode"
)

//lint:ignore ST1000 Suppressing duplicate main warning
func main() {
	input := utils.NewStdinInput()

	sum := 0
	for line := range input.Lines() {
		runes := []rune(line)

		var v int
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				vv, _ := strconv.Atoi(string(runes[i]))
				v += vv * 10
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				vv, _ := strconv.Atoi(string(runes[i]))
				v += vv
				break
			}
		}

		// log.Println(v)
		sum += v
	}

	fmt.Fprintln(os.Stderr, sum)
}
