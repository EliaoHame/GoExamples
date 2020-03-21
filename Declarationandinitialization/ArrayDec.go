package main

import (
	"fmt"
	"time"
)

var mey [40]int

func main() {
	//ArrayUse()
	//var ar [3]int
	//f(ar)
	//fp(&ar)
	//fmt.Println(ar)
	start := time.Now()
	for i := 0; i < len(mey); i++ {
		fibonacciArray(i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("memory func take time: %s\n", delta)

	start1 := time.Now()
	for i := 0; i < len(mey); i++ {
		normalFib(i)
	}
	end1 := time.Now()
	delta1 := end1.Sub(start1)
	fmt.Printf("normal func take time: %s\n", delta1)
}

func ArrayUse() {

	// one
	arr1 := [5]string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("arr1 index %d ele is: %s\n", i, arr1[i])
	}
	//two
	var arr2 [5]int
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr1 index %d ele is: %d\n", i, arr2[i])
	}

	// three
	var arr3 = new([5]int)

	arr5 := *arr3
	arr5[1] = 1000
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("arr1 index %d ele is: %d\n", i, arr5[i])
	}
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("arr1 index %d ele is: %d\n", i, arr3[i])
	}
}

func f(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	a[1] = 2
	fmt.Println(a)
}

func fibonacciArray(n int) (res int) {
	if mey[n] != 0 {
		res = mey[n]
		return
	}
	if n < 2 {
		res = 1
	} else {
		res = fibonacciArray(n-1) + fibonacciArray(n-2)
	}
	mey[n] = res
	return
}

func normalFib(n int) (res int) {
	if n <= 1 {
		res = 1
		return
	} else {
		res = normalFib(n-1) + normalFib(n-2)
	}
	return
}
