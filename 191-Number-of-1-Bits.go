package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(7))

}

func hammingWeight(num uint32) int {
	var count int
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
