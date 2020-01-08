package main

import (
	"fmt"
)

// 赋值相关
// 隐式复制
func main() {
	var x int
	x = 20
	fmt.Println(x)

	x, y := 0, 1
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)

	// 实现斐波那契数列
	fmt.Println(fib(12))

	//丢弃一部分返回值
	_, b, c := 2, 4, 6
	fmt.Println(b, c)

	// 隐式复制
	names := []string{"Bill", "Mike", "John"}
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x

}
