package main

func main() {
	for v := 1; v <= 10; v++ {
		println(sortNum(v))
	}

}

//func FibonacciOpt(n int) (res, index int) {
//	if n <= 1 {
//		res = 1
//		index = n
//	} else {
//		res1, n := FibonacciOpt(n - 1)
//		res2, n := FibonacciOpt(n - 2)
//		res = res1 + res2
//	}
//	return
//}

func sortNum(n int) int {
	if n <= 1 {
		return 10
	}
	return sortNum(n-1) - 1
}
