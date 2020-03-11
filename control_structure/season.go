package main

import "fmt"

func WhatSeason(month int) {
	switch month {
	case 3, 4, 5:
		fmt.Print("This is the spring season")
	case 6, 7, 8:
		fmt.Print("This is the summer season")
	case 9, 10, 11:
		fmt.Print("This is the autumn season")
	case 12, 1, 2:
		fmt.Print("This is the winter season")
	default:
		fmt.Print("Wrong month")
	}

}

func main() {
	WhatSeason(1)
}
