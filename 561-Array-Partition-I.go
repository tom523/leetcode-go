package main

import "sort"

func main() {
	a := []int{1, 4, 3, 2}
	arrayPairSum(a)
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ret := 0
	for index, value := range nums {
		if index%2 == 0 {
			ret += value
		}
	}
	return ret
}
