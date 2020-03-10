package main

import (
	"fmt"
	"strings"
)

func Contains(c, con string) bool {
	return strings.Contains(c, con)
}

func main() {
	par := "hello world"
	son := "hello"
	fmt.Printf("%s contains %s : %v", par, son,
		Contains(par, son))
}
