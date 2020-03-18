package main

import "fmt"

func main() {
	//g := func() {
	//	s := "hello,world"
	//	for _, content := range s {
	//		fmt.Printf("%s\n", string(content))
	//	}
	//}
	//g()
	//fmt.Printf("%T", g)
	//fmt.Println(f())
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 give: %v\n", p2(3))
	TwoAdder := Adder(2)
	// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	fmt.Printf("the result is :%v\n", TwoAdder(3))
	fmt.Printf("the result is :%v\n", TwoAdder(3))
	f := fbi()
	for i := 1; i <= 10; i++ {
		fmt.Println(f())
	}
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		a += b
		return a
	}
}

func AnonymousFunction() {
	fAdd := func(x, y int) int { return x + y }
	fmt.Println(fAdd(1, 2))
	func() {
		sum := 1
		for i := 1; i <= 100; i++ {
			sum += i
		}
	}()
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func fbi() func() int {
	b1, b2 := 0, 1
	return func() int {
		sum := b1 + b2
		b1 = b2
		b2 = sum
		return sum
	}
}
