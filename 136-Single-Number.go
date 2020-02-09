package main

import "fmt"

func main() {
	fmt.Println(singleNumber1([]int{4, 1, 2, 1, 2}))
}

// 1. delete nums[0]
// 2. find nums[0] in the rest numbers
// 3. if success, delete it
// 4. if fail, return nums[0]
// 5. repeat 1-4

func singleNumber(nums []int) int {
	curNum := nums[0]
	nums = nums[1:]
	for index, value := range nums {
		if curNum == value {
			return singleNumber(append(nums[:index], nums[index+1:]...))
		}
	}
	return curNum
}

func singleNumber1(nums []int) int {
	i := 0
	for _, value := range nums {
		i = i ^ value
	}
	return i
}
