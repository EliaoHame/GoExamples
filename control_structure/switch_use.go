package main

import "fmt"

func switchOne(i int) {
	switch i {
	case 98, 99:
		fmt.Println("It`s equal to 98")
	case 100:
		fmt.Println("It`s equal to 100")
	default:
		fmt.Println("It`s not equal to 98 or 100")
	}
}

func calculate() int {
	x, y := 10, 20
	return x + y
}

func switchTwo(i int) {
	switch {
	case i < 0:
		fmt.Printf("Number %d is a negative", i)
	case i > 0 && i <= 10:
		fmt.Printf("Number %d is between 0 and 10", i)
	default:
		fmt.Println("Number is 11 or greater")
	}
}

func switchTheer() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}

func main() {
	switchTheer()
}
