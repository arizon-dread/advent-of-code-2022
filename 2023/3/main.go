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
			r := rune(l[j])
			if r == '.' {
				continue
			}
			if unicode.IsDigit(r) {
				digit := string(r)
				//numbers[i][j] =
			out:
				for k := 1; k < len(l); k++ {
					if j+k >= len(l) {
						break out
					}
					if unicode.IsDigit(rune(l[j+k])) {
						digit += string(l[j+k])
						//fmt.Printf("another digit: %c\nnumbers: %v\n", l[j+k], numbers[i][j])
					} else {
						//increment j but subtract one to accomodate for j++
						j = j + k - 1
						break out
					}
				}
				numbers[i][index] = digit

			} else {
				//fmt.Printf("found a special char! %c\n", char)
				specialChars[i][j] = string(r)
			}
		}

	}
	for i := 0; i < len(numbers); i++ {
		for j, d := range numbers[i] {

			valid := false
			if j > 0 {
				if _, exists := specialChars[i][j-1]; exists {
					valid = true
				}
			}
			if j+len(d) >= len(specialChars[i]) {
				if _, exists := specialChars[i][j+len(d)]; exists {
					valid = true
				}
			}
			if i == 0 {
				//fmt.Printf("check below\n")
				if existsAboveOrBelow(i, j, specialChars, (len(d)), "below") {
					valid = true
				}
			} else if i == 139 {
				//fmt.Printf("check above\n")
				if existsAboveOrBelow(i, j, specialChars, (len(d)), "above") {
					valid = true
				}
			} else {
				//fmt.Printf("check both\n")
				if existsAboveOrBelow(i, j, specialChars, (len(d)), "both") {
					valid = true
				}
			}
			if valid {
				// if d != "" {
				// 	fmt.Printf("match: %v\n", d)
				// 	fmt.Printf("number: %v\n", d)
				// 	fmt.Printf("index: %v\nline: %v\n---\n", j+1, i+1)
				// }

				total += strToInt(d)
			}
		}
	}

	fmt.Printf("total: %v\n", total)
}

func strToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	fmt.Printf("error converting string to int, %v", err)
	return 0
}

func existsAboveOrBelow(row int, index int, specialChars []map[int]string, numLen int, check string) bool {
	if index == 0 {
		//make "index -1" become 0
		index = 1
	} else {
		index = index - 1
	}
	if check == "above" || check == "both" {
		for i := 0; i <= numLen+1; i++ {

			if _, exists := specialChars[row-1][index+i]; exists {
				return true
			}
		}
	}
	if check == "below" || check == "both" {
		for i := 0; i <= numLen+1; i++ {
			//fmt.Printf("index: %v\n value:", index)

			if _, exists := specialChars[row+1][index+i]; exists {
				return true
			}
		}
	}

	return false
}
