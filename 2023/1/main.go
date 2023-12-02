package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var numbers = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var n = [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

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
		//1a
		//final := getAllDigits(line)

		//1b
		final := getDigitsAndLetterNumbers(line)

		//fmt.Printf("%v and %v = %v\n", first, last, final)
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

func getAllDigits(line string) string {
	digits := strings.Map(func(r rune) rune {
		if unicode.IsNumber(r) {
			return r
		}
		return 'a'
	}, line)
	digits = strings.ReplaceAll(digits, "a", "")

	first := string(digits[0])
	last := string(digits[len(digits)-1])
	return first + last
}

func getDigitsAndLetterNumbers(line string) string {
	numberMap := map[int]int{}
	for i, v := range numbers {
		if strings.Contains(line, v) {
			index := strings.Index(line, v)
			numberMap[index] = i
			index = strings.LastIndex(line, v)
			numberMap[index] = i
		}
	}
	for i, v := range n {
		if strings.Contains(line, v) {
			index := strings.Index(line, v)
			numberMap[index] = i
			index = strings.LastIndex(line, v)
			numberMap[index] = i
		}
	}
	// fmt.Println(line)
	// print(numberMap)
	keys := make([]int, 0, len(numberMap))
	for k, _ := range numberMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	first := numberMap[keys[0]]
	last := numberMap[keys[len(keys)-1]]

	return strconv.Itoa(first) + strconv.Itoa(last)
}

func print(numberMap map[int]int) {
	for k, v := range numberMap {
		fmt.Printf("index: %v, value: %v ", k, v)
	}
	fmt.Printf("\n")
}
