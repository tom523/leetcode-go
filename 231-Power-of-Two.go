package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(0))
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	var count int
	for i := 0; i < 32; i++ {
		if n&1 == 1 {
			count++
			if count >= 2 {
				return false
			}
		}
		n = n >> 1
	}
	return true
}
