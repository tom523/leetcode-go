package main

import "strconv"

func main() {
	left := 1
	right := 22
	selfDividingNumbers(left, right)
}

func selfDividingNumbers(left int, right int) []int {
	var ret []int
	for left <= right {
		if isSelfDividing(left) {
			ret = append(ret, left)
		}
		left++
	}
	return ret
}

func isSelfDividing(i int) bool {
	for _, value := range strconv.Itoa(i) {
		str := string(value)
		divider, _ := strconv.Atoi(str)
		if divider == 0 {
			return false
		}
		if i%divider != 0 {
			return false
		}
	}
	return true
}
