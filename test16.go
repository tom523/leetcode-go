package main

import (
	"math"
	"sort"
)

func main() {
	A := []int{-4, -1, 0, 3, 10}
	sortedSquares(A)
}

func sortedSquares(A []int) []int {
	for index, value := range A {
		A[index] = int(math.Pow(float64(value), 2))
	}
	sort.Ints(A)
	return A
}

func insert(sliceA []int, insertedVal int) (ret []int) {
	for index, value := range sliceA {
		if insertedVal < value {
			b := make([]int, index, index)
			copy(b, sliceA[0:index])
			t := append(b, insertedVal)
			ret = append(t, sliceA[index:]...)
			return
		}
	}
	ret = append(sliceA, insertedVal)
	return
}
