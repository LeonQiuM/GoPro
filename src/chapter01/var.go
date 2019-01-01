package main

import (
	"fmt"
	"os"
)

// 变量
/*
var 变量名 数据类型 = 表达式
数据类型和表达式可省略其一

*/

func main() {
	var s string
	fmt.Println(s) // 字符串的默认值为空串，长度为0的字符串
	var n int
	fmt.Println(n) // int的默认值为0

	// 同时申明多个变量
	var i, j, k int
	fmt.Println(i, j, k)
	var a, b, c = true, 1, "abc"
	fmt.Println(a, b, c)

	// 函数返回多个值的接收

	var f_obj, f_err = os.Open("/Users/qiumeng/GoglandProjects/GoPro/src/chapter01/Declaration.go")
	if f_err == nil {
		fmt.Println(f_obj.Name())
		fmt.Println(f_obj.Fd())
	} else {
		fmt.Println("err")
	}
}
