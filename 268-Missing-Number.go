package main

import "fmt"

func main() {
	fmt.Println(missingNumber1([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	length := len(nums)
	sum := (length + 0) * (length + 1) / 2
	actualSum := 0
	for _, value := range nums {
		actualSum += value
	}
	return sum - actualSum
}

// XOR
func missingNumber1(nums []int) int {
	var ret int
	for index, value := range nums {
		ret ^= index ^ value
	}
	ret ^= len(nums)
	return ret
}
