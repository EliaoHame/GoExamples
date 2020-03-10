package main

import (
	"fmt"
	"math/rand"
	"time"
)

func product(header string, channel chan<- string) {
	for {
		// 随机数字格式化发送给通道
		channel <- fmt.Sprintf("%s:%v\n", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer(channel <-chan string) {
	for {
		// 通道中取出数据， 此处会阻塞直到信道中返回数据
		message := <-channel
		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建product()函数的并发 gooutine
	go product("cat", channel)
	go product("dog", channel)
	// 数据消费函数
	customer(channel)
}
