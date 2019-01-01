package main

import "fmt"

// 使用 new 创建变量
// new 是内奸函数，使用不需要导入

/*
new(T) 创建一个 T 类型的匿名变量指针
x:=new(T)
*/
func main() {
	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 123
	fmt.Println(*p)

	// 区别于普通的申明变量，不在需要变量名 ，p 为匿名变量的指针

	// new 为预定义函数，是可以被覆盖的
	var new = 20
	fmt.Println(new)
}

func NewInt1() *int {
	return new(int) // 返回 int 类型的变量的指针
}

func NewInt2() *int {
	var value int
	return &value // 返回 int 类型的变量的指针
}

//上面两个函数达到了相同的结果
