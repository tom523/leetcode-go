package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	v := x ^ y
	count := 0
	for v != 0 {
		if (v & 1) == 1 {
			count++
		}
		v = v >> 1
	}
	return count
}
