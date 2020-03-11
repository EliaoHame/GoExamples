package main

import "fmt"

func PointerCrash() {
	// 空指针反向引用会导致crash
	var p *int = nil
	*p = 0
}

func main() {
	i1 := 5
	fmt.Printf("interger: %d location in memory: %p \n", i1, &i1)

	// 声明
	var intP *int
	intP = &i1
	fmt.Printf("the value at memory location %p is %d\n", intP, *intP)

	// string pointer
	s := "good os"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("here is the pointer p: %p\n", p)
	fmt.Printf("here is a string is : %s\n", *p)
	fmt.Printf("here is a string is : %s\n", s)
	//PointerCrash()
	i := 1
	for {
		i++
		fmt.Printf("%d memory is %p \n", i, &i)
		if i == 10 {
			break
		}
	}
}
