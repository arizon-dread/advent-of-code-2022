package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculateLevels(input string) int {
	var sum int = 0
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		if uniDirectional(l) {
			if safeAdjacent(l) {
				sum++
			}
		}
	}
	return sum
}

func safeAdjacent(line string) bool {
	l := strings.Split(line, " ")
	for i, v := range l {
		if i > 0 {
			curr, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Got err!")
				return false
			}
			prev, err := strconv.Atoi(l[i-1])
			if err != nil {
				fmt.Println("Got err!")
				return false
			}
			if curr > prev && curr-prev > 3 {
				return false
			} else if curr < prev && prev-curr > 3 {
				return false
			} else if curr == prev {
				return false
			}
		}
	}
	return true
}

func uniDirectional(line string) bool {
	l := strings.Split(line, " ")
	asc := false
	desc := false
	for i, v := range l {
		if asc && desc {
			return false
		}
		if i > 0 {
			curr, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("Got err!")
				return false
			}
			prev, err := strconv.Atoi(l[i-1])
			if err != nil {
				fmt.Println("Got err!")
				return false
			}
			if curr > prev {
				asc = true
			} else if curr < prev {
				desc = true
			} else {
				return false
			}
		}
	}
	return true
}
