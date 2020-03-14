package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func greeting() {
	fmt.Println("i`m a greeting func")
}

func multReturnval(x, y int) (x1, x2 int) {
	x1 = x + y
	x2 = x - y
	return
}

func multReturnvalNoname(x, y int) (int, int) {
	sum := x + y
	minus := x - y
	return sum, minus
}

func funcType(n int, str string) string {
	x := strconv.Itoa(n) + str
	return x
}

func errorFunc(numbers float64) (nums float64, msg error) {
	if numbers < 0 {
		nums = float64(math.NaN())
		msg = errors.New("need greater 0")
	} else {
		nums = math.Sqrt(numbers)
	}
	return
}

// this function changes reply
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func myFunc(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func varargs(s ...int) {
	for _, value := range s {
		println(value)
	}
}

func main() {
	fmt.Println("I`m main before calling greeting")
	greeting()
	fmt.Println("i`m mian after calling greeting")
	x, y := errorFunc(1.21)
	fmt.Println(x, y)
	n := 0
	reply := &n
	print(*reply, "\n")
	Multiply(1, 2, reply)
	fmt.Println("n values is : ", *reply)
	fmt.Println(myFunc(1, 2, 3, 4))
	varargs(1, 3, 4, 5)
}
