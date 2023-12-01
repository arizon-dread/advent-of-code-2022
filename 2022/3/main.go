package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// var lowerCasePriority = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
// var upperCasePriority = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
func main() {
	dat, err := os.Open("input.txt")

	if err != nil {
		fmt.Printf("Error when reading file, %v\n", err)
	}
	fileScanner := bufio.NewScanner(dat)
	fileScanner.Split(bufio.ScanLines)
	var rucksacks []rucksack = make([]rucksack, 0)
	counter := 1
	var elfGroup []rune = make([]rune, 0)
	for fileScanner.Scan() {

		t := fileScanner.Text()
		middle := len(t) / 2
		sack := rucksack{}
		sack.compartments[0] = t[0:middle]
		sack.compartments[1] = t[middle:]
		calculatePriorities(&sack)
		rucksacks = append(rucksacks, sack)
		if counter == 3 {
			var groupsacks [3]rucksack
			groupsacks[0] = rucksacks[len(rucksacks)-3]
			groupsacks[1] = rucksacks[len(rucksacks)-2]
			groupsacks[2] = rucksacks[len(rucksacks)-1]
			r, err := getGroupRune(groupsacks[:])
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			elfGroup = append(elfGroup, r)
			counter = 0
		}
		counter++
		// fmt.Printf("first part: %v\n", ruckSack[0])
		// fmt.Printf("second part: %v\n", ruckSack[1])
	}
	totalPrio := 0
	for _, sack := range rucksacks {
		// if i < 10 {
		// 	fmt.Printf("compartments: %v|%v, items: %v priorities: %v\n",
		// 		sack.compartments[0], sack.compartments[1], string(sack.rearrangableItems[0]), sack.priority)
		// }
		totalPrio += sack.priority
	}
	badgePrioSum := 0
	for _, badge := range elfGroup {
		fmt.Printf("elfRune: %v\n", string(badge))
		badgePrioSum += getPriority(badge)
	}
	fmt.Printf("totalPrio (d3p1): %d\n", totalPrio)
	fmt.Printf("badgePrioSum (d3p2): %d\n", badgePrioSum)
}

func calculatePriorities(r *rucksack) {
	for _, item := range r.compartments[0] {
		for _, comparable := range r.compartments[1] {
			if item == comparable {
				//fmt.Printf("Found one! %v: %v\n", string(item), item)
				r.rearrangableItems = append(r.rearrangableItems, item)
				prio := getPriority(item)
				r.priority = prio
				//fmt.Printf("priority: %v\n", prio)
			}
		}
	}
}
func getPriority(item rune) int {

	//fmt.Printf("%v: %v, ", string(item), item)
	if item >= 97 {
		//fmt.Printf("prio: %d\n", int(item)-96)
		return int(item) - 96
	} else {
		//fmt.Printf("prio: %d\n", int(item)-38)
		return int(item) - 38
	}

}
func getGroupRune(groupsack []rucksack) (rune, error) {
	var err error = nil
	content1 := groupsack[0].compartments[0] + groupsack[0].compartments[1]
	content2 := groupsack[1].compartments[0] + groupsack[1].compartments[1]
	content3 := groupsack[2].compartments[0] + groupsack[2].compartments[1]
	for _, char1 := range content1 {
		for _, char2 := range content2 {
			if char1 == char2 {
				for _, char3 := range content3 {
					if char1 == char3 {
						fmt.Printf("%v %v %v", content1, content2, content3)
						fmt.Printf("found match: %v\n", string(char3))
						return char3, err
					}
				}
			}

		}
	}
	fmt.Printf("%v %v %v", content1, content2, content3)
	return 0, errors.New("failed to find matching badge")
}
