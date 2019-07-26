package main

import "fmt"

// 函数声明的
/*
函数的参数可以视作函数的局部变量，所以不能在相同层级定义同名的变量
任何乐行的参数传递，都是拷贝传递，，区别是拷贝对象还是拷贝指针
*/

func Jobt(x *int) {
	fmt.Printf("Pointer:,%p,target:,%v\n", &x, x)
}
func main() {
	fmt.Println("LL")
	a := 0x100
	p := &a
	fmt.Printf("Pointer:,%p,target:,%v\n", &p, p)
	//  将指针作为函数参数传入函数，收到的其实也是原指针的拷贝，内存地址不会改变
	Jobt(p)
}
