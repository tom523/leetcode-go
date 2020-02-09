package main

import (
	"strconv"
)

func main() {
	s := "10#11#12"
	freqAlphabets(s)
}

func freqAlphabets(s string) string {
	var sliceRet []byte
	var sliceT []byte
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && string(s[i+2]) == "#" {
			sliceT = append(sliceT, s[i:i+2]...)
			i += 2
		} else {
			sliceT = []uint8{s[i]}
		}
		sliceRet = append(sliceRet, numToAlphabet(sliceT))
		sliceT = nil
	}
	return string(sliceRet)
}

func numToAlphabet(sliceT []byte) byte {
	t := string(sliceT)
	a, _ := strconv.Atoi(t)
	return byte(a + 96)
}
