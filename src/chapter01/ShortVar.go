package main

import "fmt"

// 简短变量

/*
1.必须在函数内申明
2.var 申明变量，可以不初始化
3. 简短变量必须初始化
*/

func main() {
	x := 10
	y := 0.24
	fmt.Printf("x=%d,y=%f", x, y)

	a, b, c := 1, 0.2, "abc"
	fmt.Println(a, b, c)
}
