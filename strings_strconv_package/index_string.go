package main

import (
	"fmt"
	"strings"
)

func StringIndex(s, sub string) int {
	return strings.Index(s, sub)
}

func StringLastIndex(s, sub string) int {
	return strings.LastIndex(s, sub)
}

func StringRunneIndex(s string, r rune) int {
	return strings.IndexRune(s, r)
}

func main() {
	content := "Just a sample statement for testing."
	indexFind := StringIndex(content, "a")
	lastIndexFind := StringLastIndex(content, "a")
	runneIndexFind1 := StringRunneIndex(content, 101)
	runneIndexFind2 := StringRunneIndex(content, rune('e'))
	fmt.Printf("the index of `a` in content: %d \n", indexFind)
	fmt.Printf("the last index of `a` in content: %d \n", lastIndexFind)
	fmt.Printf("the last index of runne 99 in content: %d \n", runneIndexFind1)
	fmt.Printf("the last index of runne('e') in content: %d \n", runneIndexFind2)

}
