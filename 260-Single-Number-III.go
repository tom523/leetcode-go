package main

import "fmt"

func main() {
	fmt.Println(singleNumber31([]int{1, 2, 1, 3, 2, 5}))
}

// 1. delete nums[0]
// 2. find nums[0] in the rest numbers
// 3. if success, delete it
// 4. if fail, return nums[0]
// 5. repeat 1-4

func singleNumber31(nums []int) []int {
	freqMap := make(map[int]int)
	for _, value := range nums {
		freqMap[value] += 1
	}
	var ret []int
	for k, v := range freqMap {
		if v == 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
