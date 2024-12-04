package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("unable to open file, err: " + err.Error())
	}
	defer file.Close()
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("no input, reading file resulted in err: " + err.Error())
	}

	input := string(f)
	safeLevels := calculateLevels(input)
	fmt.Printf("safe number of levels %d", safeLevels)
}
