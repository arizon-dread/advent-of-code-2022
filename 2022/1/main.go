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

	var elfs []elf
	elfs = append(elfs, elf{})
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
			elfs[len(elfs)-1].snax = append(elfs[len(elfs)-1].snax, num)
		} else {
			elfs = append(elfs, elf{})
		}

	}
	var sortedElfs []elf
	for _, item := range elfs {
		calculateTotal(&item)
		sortedElfs = append(sortedElfs, item)
	}

	sort.Slice(sortedElfs, func(i, j int) bool {
		return sortedElfs[i].totalCalories < sortedElfs[j].totalCalories
	})
	fmt.Printf("total number of elves: %d \n", len(sortedElfs))
	hungriestElf := sortedElfs[len(sortedElfs)-1]
	fmt.Printf("hungriestElf calories: %d \n", hungriestElf)
	var top3elvesTotalCals int
	for i := 1; i < 4; i++ {
		top3elvesTotalCals += sortedElfs[len(sortedElfs)-i].totalCalories
	}
	fmt.Printf("Top 3 summed cals: %d", top3elvesTotalCals)

	dat.Close()
}

func calculateTotal(s *elf) {
	total := 0
	for _, item := range s.snax {
		total += item
	}
	s.totalCalories = s.totalCalories + total
}
