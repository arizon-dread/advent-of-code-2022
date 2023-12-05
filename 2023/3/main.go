package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var total int = 0

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
					if j+k > len(l) {
						break
					}
					if unicode.IsDigit(rune(l[j+k])) {
						numbers[i][j] += string(char)
					} else {
						//increment j but subtract one to accomodate for j++
						j = j + k - 1
						break nextChar
					}
				}
			} else {
				specialChars[i][j] = string(char)
			}
		}

	}
	for i, slc := range numbers {
		for j, v := range slc {
			if _, exists := specialChars[i-1][j]; exists {
				addStrToTotal(v)
			} else if _, exists := specialChars[i+1][j]; exists {
				addStrToTotal(v)
			} else if _, exists := specialChars[i][j-1]; exists {
				addStrToTotal(v)
			} else if _, exists := specialChars[i][j+len(v)+1]; exists {
				addStrToTotal(v)
			}
		}
	}
	fmt.Printf("total: %v\n", total)
}

func addStrToTotal(s string) {
	n, err := strconv.Atoi(s)
	if err == nil {
		total += n
	}
}
