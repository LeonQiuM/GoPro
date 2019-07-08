package main

import "fmt"

// 字符串
/*
不可改变的字节序列,只读
字符串通常会被解释为 utf-8编码
len() 用来获取字符串中字节的长度
可以使用索引操作：s[1]

字符串使用双引号，字符使用单引号

ascii 使用1个字节
汉字使用3个字节
*/
func main() {
	s := "Hello World"
	a := "HHHH"
	fmt.Println(len(s))
	fmt.Println(s[0], s[3]) //输出了字符的十进制ascii数值
	fmt.Println(a == "HHHH")
	// 字符串切片Slice
	s2 := "Hello 李宁" // 5+1+3+3  utf8中3个字节表示一个汉字
	fmt.Println(len(s2))
	fmt.Println(s2[6:9]) // 取头不取尾
	fmt.Println(s2[6:])  // 省略结束索引
	fmt.Println(s2[6:len(s2)])
	fmt.Println(s2[:2])
	// 字符串连接
	fmt.Println("您好" + s2[6:])
	// 字符串的比较,通过 Byte逐个比较
	ss1 := "hello"
	ss2 := "hEllo"
	fmt.Println(ss1 > ss2)
	fmt.Println(ss1 == ss2)

}
