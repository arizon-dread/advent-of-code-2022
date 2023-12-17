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

	specialChars, numbers := getCharsAndNumbers(sArr)
	total := 0

	addHitsToTotal(specialChars, numbers, &total)

	fmt.Printf("total: %v\n", total)
}

func getCharsAndNumbers(sArr []string) (map[xy]string, map[xy]string) {

	specialChars := make(map[xy]string)
	numbers := make(map[xy]string)
	// specialChars := make([]map[int]string, 140)
	// numbers := make([]map[int]string, 140)

	//loop over each line in turn
	for i, l := range sArr {
		// specialChars[i] = make(map[int]string, 10)
		// numbers[i] = make(map[int]string, 10)
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
				xy := &xy{
					x: index,
					y: i,
				}
				numbers[*xy] = digit
				//numbers[i][index] = digit

			} else {
				xy := &xy{
					x: index,
					y: i,
				}
				//fmt.Printf("found a special char! %c\n", char)
				specialChars[*xy] = string(r)
			}
		}

	}
	return specialChars, numbers
}

func addHitsToTotal(specialChars map[xy]string, numbers map[xy]string, total *int) {
	// for i := 0; i < len(numbers); i++ {
	// 	pos := &xy{
	// 		x: 0,
	// 		y: i,
	// 	}
	for k, v := range numbers {

		valid := false
		if k.x > 0 {
			if _, exists := specialChars[xy{k.x - 1, k.y}]; exists {
				valid = true
			}
		}
		//if j+len(d) <= len(specialChars[i]) && !valid {i][j+len(d)
		if _, exists := specialChars[xy{k.x + len(v), k.y}]; exists {
			valid = true
		}
		//}
		if k.y == 0 && !valid {
			//fmt.Printf("check below\n")
			if existsAboveOrBelow(k, specialChars, (len(v)), "below") {
				valid = true
			}
		} else if k.y == 139 && !valid {
			//fmt.Printf("check above\n")
			if existsAboveOrBelow(k, specialChars, (len(v)), "above") {
				valid = true
			}
		} else {
			//fmt.Printf("check both\n")
			if existsAboveOrBelow(k, specialChars, (len(v)), "both") && !valid {
				valid = true
			}
		}
		if valid {
			if v != "" {
				fmt.Printf("match: %v\n", v)
				fmt.Printf("number: %v\n", v)
				fmt.Printf("index: %v\nline: %v\n---\n", k.x+1, k.y+1)
			}

			*total += strToInt(v)
		}
	}
	//}
}

func strToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	fmt.Printf("error converting string to int, %v", err)
	return 0
}

func existsAboveOrBelow(pos xy, specialChars map[xy]string, numLen int, check string) bool {
	if pos.x == 0 {
		//make "index -1" become 0
		pos.x = 1
	} else {
		pos.x = pos.x - 1
	}
	if check == "above" || check == "both" {
		for i := 0; i <= numLen+1; i++ {

			if _, exists := specialChars[xy{pos.x + i, pos.y - 1}]; exists {
				return true
			}
		}
	}
	if check == "below" || check == "both" {
		for i := 0; i <= numLen+1; i++ {
			//fmt.Printf("index: %v\n value:", index)

			if _, exists := specialChars[xy{pos.x + i, pos.y + 1}]; exists {
				return true
			}
		}
	}

	return false
}
