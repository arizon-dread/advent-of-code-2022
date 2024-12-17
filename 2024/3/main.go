package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("could not read input")
	}
	//second star day 3:
	l := removeDonts(f)
	lines := strings.Split(string(l), "\n")

	sum := 0
	for _, line := range lines {
		sum = sum + sumFoundMuls(line)
	}
	fmt.Printf("number of instructions: %d\n", sum)
}
