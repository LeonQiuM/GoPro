package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

// 程序启动也会创建一个goroutine去执行
func main() {
	go hello() // 单独开启一个goroutine去执行这个函数，
	// 但是如果还没执行完成main结束了，由main函数启动的这个goroutine也自动结束
	fmt.Println("main")
}
