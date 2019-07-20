package main

import "fmt"

//字符串转义符
/*

 */
func main() {
	fmt.Println("ab\bx")       //  退格\b
	fmt.Print("aa\nbb")        // 换行\n
	fmt.Println("x\ry")        // 回车\r
	fmt.Println("ok\tok")      // 制表符\t
	fmt.Println("ok1\tok1")    // 制表符\t
	fmt.Println('a')           //  字符Println会打印ASCII
	fmt.Println("\" Hello \"") //  转义
	fmt.Println("\\\\\\\\")    //  转义
	fmt.Println("\x41\x42")    //  \x  16 进制
	fmt.Println("\101\102")    //  \x  8 进制
}
