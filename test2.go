package main

import "fmt"

func main() {
	indices := [][]int{
		{1, 1},
		{0, 0},
	}
	fmt.Println(oddCells(2, 2, indices))
}

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < len(indices); i++ {
		rowIndex := indices[i][0]
		columnIndex := indices[i][1]
		rowAddOne(&matrix, rowIndex)
		columnAddOne(&matrix, columnIndex)
	}
	return oddCellsCount(matrix)
}

func rowAddOne(p *[][]int, rowIndex int) {
	matrix := *p
	for c := 0; c < len(matrix[rowIndex]); c++ {
		matrix[rowIndex][c] += 1
	}
}

func columnAddOne(p *[][]int, columnIndex int) {
	matrix := *p
	for r := 0; r < len(matrix); r++ {
		matrix[r][columnIndex] += 1
	}
}

func oddCellsCount(matrix [][]int) int {
	count := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c]%2 == 1 {
				count++
			}
		}
	}
	return count
}
