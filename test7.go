package main

func main() {
	A := [][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}
	flipAndInvertImage(A)
}

func flipAndInvertImage(A [][]int) [][]int {
	lenA := len(A)
	retA := make([][]int, lenA)
	for i := 0; i < lenA; i++ {
		retA[i] = make([]int, lenA)
	}
	lenAr := len(A[0])
	for r := 0; r < lenA; r++ {
		for c := 0; c < lenAr; c++ {
			retA[r][c] = A[r][lenAr-c-1] ^ 1
		}
	}
	return retA
}
