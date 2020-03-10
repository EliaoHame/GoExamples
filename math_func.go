package main

import (
	"fmt"
	"math"
)

// int型转换int8
func uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

// float64 转换int
func intFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

func main() {
	//result, error := uint8FromInt(-1)
	//fmt.Print(result, error)
	result1 := intFromFloat64(3.55)
	fmt.Print(result1)
}
