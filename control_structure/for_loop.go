package main

import (
	"fmt"
	"strings"
)

func counter(i int) {
	end_num := i * 10
	for x := i; x < end_num; x++ {
		fmt.Printf("this is the %d iteration", x)
	}
}

func ForToString() {
	var str string = "Go is a beautiful language!"
	for x := 0; x < len(str); x++ {
		fmt.Printf("Character on position %d is :%c \n", x, str[x])
	}
}

func loop() {
	for x := 0; x < 15; x++ {
		fmt.Printf("this is the %d iteration\n", x)
	}
	i := 0
SATRT:
	fmt.Printf("the counter is at %d\n", i)
	i++
	if i < 15 {
		goto SATRT
	}
}

func ForCharacter() {
	word := "G"
	for x := 1; x <= 25; x++ {
		fmt.Printf("%s\n", strings.Repeat(word, x))
	}
}

func bitwise() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %b is %b\n", i, ^i)
	}
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		value1 := i % 3
		value2 := i % 5
		switch {
		case value1 == 0 && value2 == 0:
			fmt.Printf("FizzBuzz\n")
		case value1 == 0:
			fmt.Printf("Fizz\n")
		case value2 == 0:
			fmt.Printf("Buzz\n")
		default:
			fmt.Println(i)
		}
	}
}

func rectangle() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 20; y++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func rangestring() {
	str := "Chinese: 日本語"
	for index, value := range str {
		fmt.Printf("%d ==> %c\n", index, value)
	}
}

func for3() {
	i := 10
	for {
		i = i - 1
		fmt.Printf("the veriable i is now: %d\n", i)
		if i <= 0 {
			break
		}
	}
	fmt.Println(i)
}

func for5() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	//for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
	//	s = i+1, j+1, s+"a" {
	//	fmt.Println("Value of i, j, s:", i, j, s)
	//}
	for5()
}
