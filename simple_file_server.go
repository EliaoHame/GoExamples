package main

import "net/http"

// 创建一个简单的文件服务器
func main() {
	// 设置当前目录为http服务的根目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	// 监听8080端口
	http.ListenAndServe(":8080", nil)
}
