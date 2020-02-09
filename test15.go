package main

func main() {
	S := "IDID"
	diStringMatch(S)
}

func diStringMatch(S string) []int {
	var ret []int
	n := len(S)
	var tmp []int
	for i := 0; i <= n; i++ {
		tmp = append(tmp, i)
	}
	for _, value := range S {
		if value == 'I' {
			ret = append(ret, tmp[0])
			tmp = tmp[1:]
		} else {
			ret = append(ret, tmp[len(tmp)-1])
			tmp = tmp[0 : len(tmp)-1]
		}
	}
	return append(ret, tmp[0])
}
