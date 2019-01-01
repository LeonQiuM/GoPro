package main

import "fmt"

// 申明
/*
Go语言中的四种申明
1. var 变量
2. const 常量
3. type 类型
4. func  函数
*/
type ANC int

func main() {
	var x = 20
	const y = 20 //产量不一定需要使用
	fmt.Println(x)
	fmt.Println(y)

	var a ANC = 30 //类型
	fmt.Println(a)
}
