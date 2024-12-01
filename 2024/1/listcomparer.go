package main

import "slices"

func compareLists(a []int, b []int) int {

	slices.Sort(a)
	slices.Sort(b)
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			sum += a[i] - b[i]
		} else {
			sum += b[i] - a[i]
		}
	}
	return sum
}

func duplicateCounter(a []int, b []int) int {
	items := make([]listItem, len(a))
	for i := 0; i < len(a); i++ {
		items = append(items, listItem{i, a[i], 0})
	}
	for _, v := range b {
		for index, item := range items {
			if v == item.item {
				items[index].occurances++
			}
		}
	}
	sum := 0
	for _, v := range items {

		sum += (v.item * v.occurances)
	}
	return sum
}
