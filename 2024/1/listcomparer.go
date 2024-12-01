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
