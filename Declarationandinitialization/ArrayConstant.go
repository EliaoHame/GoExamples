package main

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
	fmt.Printf("++++++++++++++++++++++++++++++++++++++++++\n")
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	floatarray := [3]float64{7.0, 8.5, 9.1}
	fmt.Printf("float sum result: %f", ArraySum(&floatarray))
}

func ArraySum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
