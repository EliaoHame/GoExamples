package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "\u00ff\u754c"
	for k, v := range s {
		fmt.Printf("%d:%c'", k, v)
	}
	fmt.Printf("\n")
	fmt.Println(utf8.RuneCountInString(s))

	str := "hello"
	strByte := []byte(str)
	strByte[1] = 'D'
	s2 := string(strByte)
	fmt.Printf("%s", s2)
}

// 字节数组对比函数
func Compare(s1, s2 string) int {
	for i := 0; i < len(s1) && i < len(s2); i++ {
		switch {
		case s1[i] > s2[i]:
			return 1
		case s1[i] < s2[i]:
			return -1
		}

	}
	switch {
	case len(s1) > len(s2):
		return 1
	case len(s1) < len(s2):
		return -1
	}
	return 0
}
