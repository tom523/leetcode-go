package main

func main() {
	singleNumber3([]int{0, 1, 0, 1, 0, 1, 99})
}

func singleNumber3(nums []int) int {
	freqMap := make(map[int]int)
	for _, value := range nums {
		freqMap[value] += 1
	}
	var ret int
	for k, v := range freqMap {
		if v == 1 {
			ret = k
			break
		}
	}
	return ret
}
