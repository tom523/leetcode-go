package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}

func uniqueMorseRepresentations(words []string) int {
	var morseMap map[string]int
	morseMap = make(map[string]int)
	for _, value := range words {
		morse := convertWordToMorse(value)
		morseMap[morse] = 0
	}
	return len(morseMap)
}

func convertWordToMorse(s string) string {
	morseTable := make(map[int32]string)
	morseArray := [26]string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	for index, value := range morseArray {
		morseTable['a'+int32(index)] = value
	}
	var ret string
	for _, value := range s {
		ret = strings.Join([]string{ret, morseTable[value]}, "")
	}
	return ret
}

func isInSlice(sliceS []string, str string) bool {
	for _, value := range sliceS {
		if value == str {
			return true
		}
		continue
	}
	return false
}
