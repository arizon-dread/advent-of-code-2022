package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic("no input, reading file resulted in err: " + err.Error())
	}

	input := string(f)
	safeLevels := calculateLevels(input)
	fmt.Printf("safe number of levels %d", safeLevels)
}
