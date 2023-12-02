package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("error reading input file , %v", err)
		os.Exit(1)
	}

	s := string(f)
	sArr := strings.Split(s, "\n")

	validRed := 12
	validGreen := 13
	validBlue := 14

	total := 0

	for i, l := range sArr {
		if l == "" {
			continue
		}

		reg := regexp.MustCompile(`^Game\s\d*:\s`)
		game := reg.ReplaceAllString(l, "")
		//fmt.Printf("%v\n", t)
		sets := strings.Split(game, ";")
		gameIsValid := true
		for _, set := range sets {
			var r int
			var g int
			var b int
			fmt.Println()
			cubes := strings.Split(set, ",")
			for _, c := range cubes {
				c = strings.TrimSpace(c)
				numCol := strings.Split(c, " ")
				fmt.Printf("number and color: %v, %v\n", numCol[0], numCol[1])
				switch numCol[1] {
				case "red":
					addNum(&r, numCol[0])
				case "green":
					addNum(&g, numCol[0])
				case "blue":
					addNum(&b, numCol[0])
				}
			}

			if r <= validRed && g <= validGreen && b <= validBlue {
				fmt.Printf("valid: %v\nred: %v, green: %v, blue: %v\n", i, r, g, b)

			} else {
				fmt.Printf("invalid: %v\nred: %v, green: %v, blue: %v\n", i, r, g, b)
				gameIsValid = false
			}
		}
		if gameIsValid {
			total += i
		}

	}
	fmt.Printf("total: %v", total)
}

func addNum(col *int, num string) {
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Printf("unable to cast %v to int: %v\n", num, err)
		return
	}
	*col += n
}
