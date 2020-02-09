package main

import "strings"

func main() {
	findTheDifference("a", "aa")
}

func findTheDifference(s string, t string) byte {
	var ret byte
	for _, value := range t {
		if !strings.ContainsRune(s, value) {
			ret = byte(value)
		} else {
			s = strings.Replace(s, string(value), "", 1)
		}
	}
	return ret
}
