package main

import (
	"fmt"
)

// 指针在函数中的应用
/*
1. 函数返回指针
2. 指针作为函数的参数


*/

func main() {
	/*Go 语言中，返回局部变量的指针，是安全的*/
	// 复制指针，为原变量创建别名

	var xp = ff()
	fmt.Println(xp)
	fmt.Println(ff() == ff()) // 每次调用都分配到不同的内存空间
	fmt.Println(*xp)

	// 将指针类型作为函数参数
	value := 10
	ChangeValue := fun1(value)
	fmt.Println(value) // 这个输出的值还是为10,等于修改的是 value 的副本
	fmt.Println(ChangeValue)

	ChangeValue2 := fun2(&value)
	fmt.Println(value) // 此处fun2传入的值为指针的副本，函数中直接修改了指针内存空间对应的变量
	// 意思为，多个内存地址指向了同一个变量，
	fmt.Println(ChangeValue2)

}

func ff() *int {
	// 返回一个指向一个int类型的指针
	v := 10
	return &v
}

func fun1(value int) int {
	value = value * 2
	return value

}

func fun2(value *int) int {
	*value = *value * 2
	return *value

}
