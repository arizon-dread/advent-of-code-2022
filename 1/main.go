package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	dat, err := os.Open("input.txt")

	var content []snack
	content = append(content, snack{})
	if err != nil {
		fmt.Printf("shit went south, %v\n", err)
	}
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split((bufio.ScanLines))
	fmt.Println("read file")
	for fileScanner.Scan() {
		t := fileScanner.Text()

		num, err := strconv.Atoi(t)
		if err == nil {
			//fmt.Printf("line: %v", num)
			content[len(content)-1].snax = append(content[len(content)-1].snax, num)
		} else {
			content = append(content, snack{})
		}

	}
	var result []snack
	for _, item := range content {
		calculateTotal(&item)
		result = append(result, item)
	}
	//fmt.Printf("does this have calories: %d ? \n", result[0].totalCalories)

	sort.Slice(result, func(i, j int) bool {
		return result[i].totalCalories < result[j].totalCalories
	})
	fmt.Printf("total number of elves: %d \n", len(result))
	snarre := result[len(result)-1]
	fmt.Printf("snarre calories: %d \n", snarre)

	dat.Close()
}

func calculateTotal(s *snack) {
	total := 0
	for _, item := range s.snax {
		total += item
	}
	s.totalCalories = s.totalCalories + total
	//fmt.Printf("does this have calories: %d ? \n", s.totalCalories)
}
