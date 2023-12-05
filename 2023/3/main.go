package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("error reading input file, %v", err)
		os.Exit(1)
	}
	s := string(f)

	sArr := strings.Split(s, "\n")

	specialChars := make([]map[int]string, 140)
	numbers := make([]map[int]string, 140)
	var total int = 0
	//loop over each line in turn
	for i, l := range sArr {
		specialChars[i] = make(map[int]string, 10)
		numbers[i] = make(map[int]string, 10)
		//loop over each char on the line
		for j := 0; j < len(l); j++ {
			char := rune(l[j])
			if char == '.' {
				continue
			}
			if unicode.IsDigit(char) {
				numbers[i][j] = string(char)
			nextChar:
				for k := 1; k < len(l); k++ {
					if j+k >= len(l) {
						break
					}
					if unicode.IsDigit(rune(l[j+k])) {
						numbers[i][j] += string(l[j+k])
					} else {
						//increment j but subtract one to accomodate for j++
						j = j + k - 1
						break nextChar
					}
				}
			} else {
				fmt.Printf("found a special char! %c\n", char)
				specialChars[i][j] = string(char)
			}
		}

	}
	for i, slc := range numbers {

		for j, v := range slc {
			//fmt.Printf("found a number! %v\n", v)
			valid := false
			if i > 0 && j > 0 && len(specialChars[i]) >= j+len(v)+1 {
				if _, exists := specialChars[i-1][j]; exists {
					valid = true
				}
				if _, exists := specialChars[i+1][j]; exists {
					valid = true
				}
				if _, exists := specialChars[i][j-1]; exists {
					valid = true
				}
				if _, exists := specialChars[i][j+len(v)+1]; exists {
					valid = true
				}
			}
			if valid {
				total += addStrToTotal(v)
			}
		}
	}
	fmt.Printf("total: %v\n", total)
}

func addStrToTotal(s string) int {
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	return 0
}
