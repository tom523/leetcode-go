package main

import "fmt"

func main() {
	S := "(()())(())"
	lcount := 0
	rcount := 0
	var ret []byte
	var tempS []byte
	for _, value := range S {
		if string(value) == "(" {
			lcount += 1
		} else {
			rcount += 1
		}
		if lcount != 1 && lcount != rcount {
			tempS = append(tempS, byte(value))
		}
		if lcount == rcount {
			ret = append(ret, tempS...)
			lcount = 0
			rcount = 0
			tempS = nil
		}
	}
	fmt.Println(string(ret[:]))
}
