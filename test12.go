package main

import "fmt"

func main() {
	moves := "UDU"
	fmt.Println(judgeCircle(moves))
}

func judgeCircle(moves string) bool {
	Rcount := 0
	Lcount := 0
	Ucount := 0
	Dcount := 0
	for _, value := range moves {
		switch value {
		case 'R':
			Rcount++
		case 'L':
			Lcount++
		case 'U':
			Ucount++
		case 'D':
			Dcount++
		}
	}
	return Rcount == Lcount && Ucount == Dcount
}
