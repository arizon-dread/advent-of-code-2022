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
			index := j
			fmt.Printf("index: %v\nline: %v\n---\n", index, i)
			char := rune(l[j])
			if char == '.' {
				continue
			}
			if unicode.IsDigit(char) {
				numbers[i][j] = string(char)
			out:
				for k := 1; k < len(l); k++ {
					if j+k >= len(l) {
						break out
					}
					if unicode.IsDigit(rune(l[j+k])) {
						numbers[i][j] += string(l[j+k])
						fmt.Printf("another digit: %c\nnumbers: %v\n", l[j+k], numbers[i][j])
					} else {
						//increment j but subtract one to accomodate for j++
						j = j + k - 1
						break out
					}
				}
				digit := numbers[i][j]
				fmt.Printf("number: %v\n", digit)
				valid := false
				if index > 0 {
					if _, exists := specialChars[i][index-1]; exists {
						valid = true
					}
				}
				if len(specialChars[i]) >= index+len(digit)+1 {
					if _, exists := specialChars[i][index+len(digit)+1]; exists {
						valid = true
					}
				}
				if i == 0 {
					fmt.Printf("check below\n")
					if existsAboveOrBelow(i, index, specialChars, (len(digit)), "below") {
						valid = true
					}
				} else if i == 140 {
					fmt.Printf("check above\n")
					if existsAboveOrBelow(i, index, specialChars, (len(digit)), "above") {
						valid = true
					}
				} else {
					fmt.Printf("check both\n")
					if existsAboveOrBelow(i, index, specialChars, (len(digit)), "both") {
						valid = true
					}
				}
				if valid {
					total += strToInt(digit)
				}
			} else {
				fmt.Printf("found a special char! %c\n", char)
				specialChars[i][j] = string(char)
			}
		}

	}

	fmt.Printf("total: %v\n", total)
}

func strToInt(s string) int {
	fmt.Printf("Convertin %v to int\n", s)
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	return 0
}

func existsAboveOrBelow(row int, index int, specialChars []map[int]string, numLen int, check string) bool {
	if index == 0 {
		//make "index -1" become 0
		index = 1
	}
	if check == "above" || check == "both" {
		for i := index - 1; i >= numLen+1; i++ {
			if _, exists := specialChars[row-1][i]; exists {
				return true
			}
		}
	}
	if check == "below" || check == "both" {
		for i := index - 1; i >= numLen+1; i++ {
			if _, exists := specialChars[row+1][i]; exists {
				return true
			}
		}
	}

	return false
}
