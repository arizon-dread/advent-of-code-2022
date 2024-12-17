package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func mul(x int, y int) int {
	return x * y
}

func sumFoundMuls(line string) int {
	regex := regexp.MustCompile(`mul\([0-9]{1,3}\,[0-9]{1,3}\)`)
	hits := regex.FindAll([]byte(line), -1)
	sum := 0
	for _, v := range hits {
		str := string(v)
		comma := strings.Index(str, ",")
		pre, err := strconv.Atoi(str[4:comma])
		if err != nil {
			fmt.Printf("could not cast pre to int, %v\n", str[4:comma])
			continue
		}
		post, err := strconv.Atoi(str[comma+1 : len(str)-1])
		if err != nil {
			fmt.Printf("could not cast post to int, %v\n", str[comma+1:len(str)-1])
			continue
		}
		sum = sum + mul(pre, post)
	}
	return sum
}
func removeDonts(input []byte) []byte {
	regex := regexp.MustCompile(`(?U)don\'t\(\)(.*\n?)*do\(\)`)
	return regex.ReplaceAll(input, []byte{})

}
