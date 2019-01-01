package main

import "fmt"

// 指针

/*
存储的值为内存地址
1. 任何类型的指针变量的的零值是 nil，也就是空指针，不指向内存地址: p = nil

两个指针相等：两个指针指向了同一个内存地址，或都为 nil

*/

func main() {
	x := 12
	fmt.Println(x)
	p := &x // 获取x变量的内存地址
	fmt.Println(p)
	fmt.Println(*p) // 获取指针变量的对应的值

	*p = 32
	fmt.Println(x) // 将 p 指针对应的值赋值改写为32

	// 指针为空
	var x1, y1 int
	fmt.Println(&x1 == &x1, &x1 == &y1, &x1 == nil)
	fmt.Println(&x1) // 虽然没有为变量 x1赋值，但是程序任然为 x1 分配了内存地址

}
