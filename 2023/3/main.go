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

	specialChars, numbers, potentialGears := getCharsAndNumbersAndPotentialGears(sArr)
	total := 0
	gearRatio := 0

	addHitsToTotal(specialChars, &potentialGears, numbers, &total)

	calcGearRatio(potentialGears, &gearRatio)

	fmt.Printf("total: %v\n", total)
	fmt.Printf("gearRatio: %v\n", gearRatio)
}

func calcGearRatio(potentialGears map[xy]string, gearRatio *int) {
	for _, v := range potentialGears {
		gear := strings.TrimPrefix(v, "*")
		gear = strings.TrimSuffix(gear, "*")
		tmpGears := strings.Split(gear, "*")
		var gears []string
		if len(tmpGears) > 2 {
			gears = removeDuplicateStr(tmpGears)
		} else {
			gears = tmpGears
		}

		if len(gears) == 2 {
			gear1, err1 := strconv.Atoi(gears[0])
			gear2, err2 := strconv.Atoi(gears[1])
			if err1 == nil && err2 == nil {
				*gearRatio += (gear1 * gear2)
				//fmt.Printf("gear at row: %v, col: %v, gear1: %v, gear2: %v\n", k.y+1, k.x+1, gear1, gear2)
				// if gear1 == gear2 {
				// 	fmt.Printf("DUPLICATE: pos: y: %v, x: %v, gear1: %v, gear2: %v \n", k.y, k.x, gear1, gear2)
				// }
			}
			//too high: 84836005
			//correct: 84159075
			//too low: 84084546
		} else {
			//fmt.Printf("potential gear with length != 2 at row: %v, col: %v. length was: %v, gears were: %v\n", k.y+1, k.x+1, len(gears), gears)
		}
	}
}
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func getCharsAndNumbersAndPotentialGears(sArr []string) (map[xy]string, map[xy]string, map[xy]string) {

	specialChars := make(map[xy]string)
	numbers := make(map[xy]string)
	potentialGears := make(map[xy]string)

	//loop over each line in turn
	for i, l := range sArr {

		//loop over each char on the line
		for j := 0; j < len(l); j++ {
			index := j
			r := rune(l[j])
			if r == '.' {
				continue
			}
			if unicode.IsDigit(r) {
				digit := string(r)
			out:
				for k := 1; k < len(l); k++ {
					if j+k >= len(l) {
						break out
					}
					if unicode.IsDigit(rune(l[j+k])) {
						digit += string(l[j+k])
					} else {
						//increment j, subtract one to accomodate for j++
						j = j + k - 1
						break out
					}
				}
				xy := &xy{
					x: index,
					y: i,
				}
				numbers[*xy] = digit

			} else {
				xy := &xy{
					x: index,
					y: i,
				}
				specialChars[*xy] = string(r)
				if r == '*' {
					potentialGears[*xy] = string(r)
				}
			}
		}

	}
	return specialChars, numbers, potentialGears
}

func addHitsToTotal(specialChars map[xy]string, potentialGears *map[xy]string, numbers map[xy]string, total *int) {

	for k, v := range numbers {

		valid := false
		if k.x > 0 {
			if _, exists := specialChars[xy{k.x - 1, k.y}]; exists {
				valid = true
			}
			if c, exists := (*potentialGears)[xy{k.x - 1, k.y}]; exists {
				c += fmt.Sprintf("%v*", v)
				(*potentialGears)[xy{k.x - 1, k.y}] = c
			}
		}
		if _, exists := specialChars[xy{k.x + len(v), k.y}]; exists {
			valid = true
		}
		if c, exists := (*potentialGears)[xy{k.x + len(v), k.y}]; exists {
			c += fmt.Sprintf("%v*", v)
			(*potentialGears)[xy{k.x + len(v), k.y}] = c
		}
		//}
		if k.y == 0 {
			if exists, _ := existsAboveOrBelow(k, specialChars, len(v), "below"); exists {
				valid = true
			}
			if exists, pos := existsAboveOrBelow(k, (*potentialGears), len(v), "below"); exists {
				c := (*potentialGears)[pos]
				c += fmt.Sprintf("%v*", v)
				(*potentialGears)[pos] = c
			}
		} else if k.y == 139 {
			if exists, _ := existsAboveOrBelow(k, specialChars, len(v), "above"); exists {
				valid = true
			}
			if exists, pos := existsAboveOrBelow(k, (*potentialGears), len(v), "above"); exists {
				c := (*potentialGears)[pos]
				c += fmt.Sprintf("%v*", v)
				(*potentialGears)[pos] = c
			}
		} else {
			if exists, _ := existsAboveOrBelow(k, specialChars, (len(v)), "both"); exists {
				valid = true
			}
			if exists, pos := existsAboveOrBelow(k, (*potentialGears), len(v), "both"); exists {
				c := (*potentialGears)[pos]
				c += fmt.Sprintf("%v*", v)
				(*potentialGears)[pos] = c
			}
		}
		if valid {
			*total += strToInt(v)
		}
	}
}

func strToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	fmt.Printf("error converting string to int, %v", err)
	return 0
}

func existsAboveOrBelow(pos xy, specialChars map[xy]string, numLen int, check string) (bool, xy) {
	if pos.x != 0 {
		pos.x = pos.x - 1
	}
	if check == "above" || check == "both" {
		for i := 0; i <= numLen+1; i++ {

			if _, exists := specialChars[xy{pos.x + i, pos.y - 1}]; exists {
				return true, xy{pos.x + i, pos.y - 1}
			}
		}
	}
	if check == "below" || check == "both" {
		for i := 0; i <= numLen+1; i++ {

			if _, exists := specialChars[xy{pos.x + i, pos.y + 1}]; exists {
				return true, xy{pos.x + i, pos.y + 1}
			}
		}
	}

	return false, xy{}
}
