package main

import "fmt"

func main() {
	//var slice []int = make([]int, 4, 18)
	//for i := 0; i < len(slice); i++ {
	//	slice[i] = i + 1
	//}
	//
	//for ix, value := range slice {
	//	fmt.Printf("Slice at %d value is %d\n", ix, value)
	//}
	//
	//items := [...]int{10, 20, 30, 40, 50}
	//for _, item := range items {
	//	item *= 2
	//	println(item)
	//}
	//CreateSlice()
	//copy_append()
	var slice []int = make([]int, 4, 18)
	sliceExtend(slice, 2)
}

func CreateSlice() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	println(len(slice1[1:1]))

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

func copy_append() {
	slF := []int{1, 2, 3}
	slT := make([]int, 10)
	n := copy(slT, slF)
	fmt.Println(slT)
	fmt.Printf("copy %d ele\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

func sliceExtend(s []int, factor int) {
	sLen := len(s)
	extendLen := sLen * factor
	if extendLen > cap(s) {
		NewSlice := make([]int, extendLen)
		s = NewSlice
	}
	s = s[0:extendLen]
	println(len(s))
}
