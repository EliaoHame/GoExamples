package main

import (
	"fmt"
	"strings"
)

func StringPrefix(s, per string) bool {
	return strings.HasPrefix(s, per)
}

func StringSuffix(s, per string) bool {
	return strings.HasSuffix(s, per)
}

func main() {
	str := "This is Go programmingT"
	ch := "T"
	pre := StringPrefix(str, ch)
	suf := StringSuffix(str, ch)
	fmt.Printf("%s start with %s: %v \n", str, ch, pre)
	fmt.Printf("%s end with %s: %v", str, ch, suf)
}
