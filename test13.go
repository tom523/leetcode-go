package main

import "sort"

func main() {
	heights := []int{1, 1, 4, 2, 1, 3}
	heightChecker(heights)
}

func heightChecker(heights []int) int {
	ret := 0
	heightsCopy := make([]int, len(heights))
	copy(heightsCopy, heights)
	sort.Ints(heights)
	for i := 0; i < len(heights); i++ {
		if heightsCopy[i] != heights[i] {
			ret++
		}
	}
	return ret
}
