package main

import "fmt"

func labelUse() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is %d, and j is %d\n", i, j)
		}
	}
}

func gotoUse() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func main() {
	gotoUse()
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
