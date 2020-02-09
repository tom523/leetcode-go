package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := '5'
	fmt.Printf("%T", reflect.TypeOf(x))
}
