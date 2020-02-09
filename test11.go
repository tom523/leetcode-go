package main

import "fmt"

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}

func replaceElements(arr []int) []int {
	for index, _ := range arr {
		if len(arr)-index == 1 {
			arr[index] = -1
		} else {
			arr[index] = maxValue(arr[index+1:])
		}
	}
	return arr
}

func maxValue(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if max < value {
			max = value
		}
	}
	return max
}
