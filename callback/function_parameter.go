package main

import (
	"fmt"
	"strings"
)

func main() {
	//callback(1, Add)
	fmt.Println(IsAscii(12))
	s := "this is a viedo 测试测试 hahahah"
	fmt.Printf(strings.Map(IsAscii, s))
}

func Add(a, b int) {
	fmt.Printf("the sun of %d and %d is : %d \n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func IsAscii(r rune) rune {
	if r > 255 {
		return '?'
	}
	return r
}
