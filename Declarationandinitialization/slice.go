package main

import (
	"bytes"
	"fmt"
)

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	var sl []byte
	Append(sl, []byte{'g', 'o'})

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("the slice1 max len is %d\n", cap(slice1))
	ByteSlice()
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(SliceSum(arr[2:5]))
	Example()
	BufferExample()
}

func SliceSum(a []int) int {
	v := 0
	for i := 0; i < len(a); i++ {
		v += a[i]
	}
	return v
}

func MakeSlice() {
	// make创建切片
	marry := make([]int, 10)
	var marry1 []int = make([]int, 10)
	fmt.Println(marry, marry1)

	//两种生成相同切片的方式
	v1 := make([]int, 50, 100)
	var v2 []int = new([100]int)[0:50]
	fmt.Printf("%d======%d", v1, v2)
}

func ByteSlice() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("%s\n", b[1:4])
	fmt.Printf("%s\n", b[:2])
	fmt.Printf("%s\n", b[2:])
	fmt.Printf("%s\n", b[:])

}

func Example() {
	s := make([]byte, 5)
	fmt.Println(len(s), cap(s))
	s = s[2:4]
	fmt.Println(len(s), cap(s))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("%s ========= %s\n", s1, s2)
	s2[1] = 't'
	fmt.Printf("%s ========= %s\n", s1, s2)
}

func BufferExample() {
	var buffer bytes.Buffer
	content := []string{"h", "e", "l", "l", "o"}
	for i := 0; i < len(content); i++ {
		buffer.WriteString(content[i])
	}
	fmt.Print(buffer.String())
}

func Append(slice, data []byte) {
	for i := 0; i < len(data); i++ {
		slice = append(slice, data[i])
	}
	println("====================\n")
	println(&slice)
	println("====================\n")
}

func InsertStringSlice(s1, s2 []string, index int) {
	new_slice := make([]string, len(s1)+len(s2))
	if index <= len(new_slice) {
		at := copy(new_slice, s2[:index])
		at += copy(new_slice[at:], s1)
		copy(new_slice[at:], s2[index:])
		fmt.Println(new_slice)
	}
	panic("error index value")
}

func RemoveStringSlice(StringList []string, start, end int) {
	if end > start {
		NewString := make([]string, len(StringList)-(end-start))
		index := copy(NewString, NewString[:start])
		copy(NewString[index:], NewString[end:])
		fmt.Println(NewString)
	}

}

func sliceExtend(s []int, factor int) {
	sLen := len(s)
	extendLen := sLen * factor
	if extendLen > cap(s) {
		NewSlice := make([]int, extendLen)
		s = NewSlice
	}
	s = s[0:extendLen]
}

func CopyAppend() {
	slF := []int{1, 2, 3}
	slT := make([]int, 10)
	n := copy(slT, slF)
	fmt.Println(slT)
	fmt.Printf("copy %d ele\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
