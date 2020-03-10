package main

import (
	"fmt"
	"strings"
)

func ReplaceStr(str, new, old string, n int) string {
	return strings.Replace(str, new, old, n)
}

func RepeatStr(str string, counts int) string {
	return strings.Repeat(str, counts)
}

func main() {
	content := "the way to go"
	fmt.Printf("replace use result: %s \n",
		ReplaceStr(content, "o", "fuck", -1))

	fmt.Printf("repeat us : %s", RepeatStr(content, 3))
}
