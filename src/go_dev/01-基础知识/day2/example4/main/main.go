package main

import "fmt"

func modify(a int) {
	a = 10
	return
}

func modifyY(a *int) {
	*a = 10
	return
}

func main() {
	a := 5                 //  值类型，内存通常是在栈中分配
	b := make(chan int, 5) // 引用类型，内存通常在堆中分配，通过GC来回收
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	//
	modify(a) // 值类型用作形参传入，等于 copy 了一份，modify 修改的 copy 的那一份
	fmt.Println(a)
	//
	modifyY(&a) //引用类型用作形参传入，传入的为其引用的地址，修改仍单修改的为原值
	fmt.Println(a)

}
