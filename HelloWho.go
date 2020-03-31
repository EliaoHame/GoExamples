package main

import (
"fmt"
"os"
"strings"
)

var big = [][]string{
	{"0", "0 0", "0  0", "0 0", "0"},
	{"11", "1", "1", "1", "111"},
	{"2", "22", "222", "   2", "   2222"},
}

func main() {
	who := "world"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("hello1", who)
}
