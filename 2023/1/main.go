package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

//var numberPresent = map[string]bool{"one": true, "two": true, "three": true, "four": true, "five": true, "six": true, "seven": true, "eight": true, "nine": true}

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input: %v\n", err)
		os.Exit(1)
	}
	s := string(f)
	sArr := strings.Split(s, "\n")
	total := 0
	for _, line := range sArr {
		// numberMap := map[int]int{}
		// for i, v := range numbers {
		// 	if strings.Contains(line, v) {
		// 		index := strings.Index(line, v)
		// 		numberMap[index] = i
		// 	}
		// }

		digits := strings.Map(func(r rune) rune {
			if unicode.IsNumber(r) {
				return r
			}
			return 'a'
		}, line)
		digits = strings.ReplaceAll(digits, "a", "")

		first := string(digits[0])
		last := string(digits[len(digits)-1])

		final := first + last
		n, err := strconv.Atoi(final)
		if err != nil {
			fmt.Printf("error converting string %v to int: %v", final, err)
			os.Exit(1)
		}
		total += n
		// fmt.Printf("%v\n", line)
	}
	fmt.Printf("total: %v\n", total)
}

func replaceWithDigit(s string) (int, error) {
	for i, v := range numbers {
		if s == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("error finding index of %v", s)
}
