package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var orig string = "ABV"
	var newS string

	fmt.Printf("size of ints is %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("error has happend %s", err)
		os.Exit(1)
	}
	fmt.Printf("the integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
