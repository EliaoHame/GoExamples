package main

import (
	"fmt"
	"os"
	"runtime"
)

// 获取当前操作系统和环境变量
func main() {
	var goos string = runtime.GOOS
	if goos == "windows" {
		fmt.Printf("The operating system is: %s\n", goos)
		path := os.Getenv("PATH")
		fmt.Printf("Path is %s\n", path)
	} else {
		println("Unknow System")
	}

}
