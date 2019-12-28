package main

import (
	"fmt"
)

func main() {
	feb(10)
}

func feb(num int) {
	// 定义一个切片并初始化
	var a []int
	a = make([]int, num)

	a[0] = 1
	a[1] = 1

	for i := 2; i < num; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	for _, i := range a {
		fmt.Println(i)
	}
}
