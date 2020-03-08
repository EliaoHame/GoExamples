package main

import (
	"fmt"
	"strconv"
)

func IntSizeUse() int {
	return strconv.IntSize
}

func AtoiUse(s string) (i int, err error) {
	return strconv.Atoi(s)
}

func main() {
	fmt.Printf("The pc sieze of int is %d \n", IntSizeUse())
	an, _ := AtoiUse("12345678")
	fmt.Printf("string number to int is :%d", an)
}
