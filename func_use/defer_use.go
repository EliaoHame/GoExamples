package main

import "fmt"

func function1() {
	fmt.Println("function1 called")
	defer function2()
	fmt.Println("function1 end called")
}

func function2() {
	fmt.Println("function1 after runing i called")
}

func deferUse() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}

// deger simulate database discounnected
func connectDb() {
	fmt.Println("connect db success")
}

func disconnecteDb() {
	fmt.Println("database discounnected")
}

func main() {
	connectDb()
	defer disconnecteDb()
	fmt.Println("do something after connecte database")
	fmt.Println("return database result")
	return
}
