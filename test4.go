package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(-123))
}

func reverse(x int) int {
	strX := strconv.Itoa(x)
	if x < 0 {
		strX = strX[1:]
	}
	var ret int
	for index, value := range strX {
		strValue := string(value)
		intValue, _ := strconv.Atoi(strValue)
		ret += int(math.Pow10(index)) * intValue
	}
	if x < 0 {
		ret = -ret
	}
	if float64(ret) < -math.Pow(2, 31) || float64(ret) > math.Pow(2, 31) {
		return 0
	}
	return ret
}
