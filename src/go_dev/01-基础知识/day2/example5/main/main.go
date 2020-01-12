package main

import "fmt"

var a string
var x1 int = 3333333333333333333
var x2 int8 = 127
var x3 int16 = 255
var x4 int32 = 32767
var x5 int64 = 3333333333333333333

func main() {
	var x = 1
	var y = 2
	fmt.Printf("x=%d,y=%d\n", x, y)
	changePoint(&x, &y)
	a = "xxx"
	fmt.Println(a)
	f1()

	//  fmt.Printf("二进制:%d，8进制:%[1]o, 16进制:%[1]x, 0x的16进制:%#[1]x", oo)
	fmt.Printf("%[1]b\n", x2)
	fmt.Println(x2 + 1) // -128:  01111111  + 1 = 10000000   首位 1 表示负号
}

func f1() {
	a := "sss"
	fmt.Println(a)
	f2() // f2 中打印的 a，由于f2 中并没有定义，且f1函数的执行已经切换，故 f2 中会使用全局变量的 a
}
func f2() {
	fmt.Println(a)
}

func changePoint(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
	fmt.Printf("x=%d,y=%d\n", *a, *b)
}
