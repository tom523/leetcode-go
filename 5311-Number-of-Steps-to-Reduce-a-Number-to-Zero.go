package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(8))
}

var count int

func numberOfSteps(num int) int {
	var count int
	for num != 0 {
		count++
		if num%2 == 0 {
			num = num / 2
		} else {
			num = num - 1
		}
	}
	return count
}
