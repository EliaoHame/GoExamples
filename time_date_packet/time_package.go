package main

import (
	"fmt"
	"time"
)

func useTime() {
	// 当前时间
	today := time.Now()
	fmt.Printf("today is : %4d-%02d-%02d %02d:%02d:%02d \n", today.Year(),
		today.Month(), today.Day(), today.Hour(), today.Minute(), today.Second())

	// 通用协调世界时间
	t := time.Now()
	fmt.Printf("UTC time : ")
	fmt.Println(t)

	week := 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)

	//time.After() time.Ticker

	time.Sleep(10) // 实现某个进程时长为d的暂停
	// 预定义格式时间
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))

}

func main() {
	useTime()
}
