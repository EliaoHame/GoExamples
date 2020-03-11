package main

import (
	"fmt"
	"runtime"
	"strconv"
)

var prompt = "Enter a digit, e.g.3" + " or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D, Enter")
	}
}

func guessNumber(n int) {
	guessNumber := 85
	if n > guessNumber {
		println("too more")
	} else if n < guessNumber {
		println("too less")
	} else {
		println("bingo")

	}

}

func ABs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

func isGreater(x, y int) bool {
	if x < y {
		return true
	}
	return false
}

func main() {
	anInt, er := strconv.Atoi("1s231")
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(anInt)
}
