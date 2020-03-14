package main

import "fmt"

func main() {
	fmt.Println(jumpStep(10))
	result := 10
	for i := 1; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 阶乘
func jumpStep(n int) (res int) {
	if n <= 2 {
		res = 2
	} else {
		res = n * jumpStep(n-1)
	}
	return
}
