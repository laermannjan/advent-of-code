package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func extractDigitFromWord(word string) int {
	digit_names := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(digit_names); i++ {
		if strings.HasPrefix(word, digit_names[i]) {
			return i
		}
	}
	return -1
}

func a() {
	file, err := os.Open(os.Getenv("AOC_DATA_ROOT") + "/2023/examples/01-a.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
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

		fmt.Println(v)
		sum += v
	}
	fmt.Println(sum)
}

func b() {
	file, err := os.Open(os.Getenv("AOC_DATA_ROOT") + "/2023/examples/01-b.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	digit_names := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		var firstDigit int
		var secondDigit int

	left_letter_loop:
		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				firstDigit, _ = strconv.Atoi(string(runes[i]))
				break
			} else {
				for d, name := range digit_names {
					if strings.HasPrefix(string(runes[i:]), name) {
						firstDigit = d
						break left_letter_loop
					}
				}
			}
		}

	right_letter_loop:
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				secondDigit, _ = strconv.Atoi(string(runes[i]))
				break
			} else {
				for d, name := range digit_names {
					if strings.HasSuffix(string(runes[:i+1]), name) {
						secondDigit = d
						break right_letter_loop
					}
				}
			}
		}

		fmt.Println(line)
		fmt.Println(firstDigit, secondDigit)
		sum += firstDigit*10 + secondDigit
	}
	fmt.Println(sum)
}

func main() {
	// a()
	b()
}
