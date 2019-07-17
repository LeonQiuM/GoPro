package main

import (
	"fmt"
	"reflect"
)

// 申明
/*
Go语言中的四种申明
1. var 变量
2. const 常量，常量表示程序运行时恒定不变的值,可以显式的置顶类型，指定类型后的复制需要前后一致
	—
3. type 类型
4. func  函数
*/
type ANC int

const (
	i, j, k                      = 1, 2, 0.1
	b                            = false
	n       int64                = 210
	m                            //  常量组中如果不指定类型和初始化值，那么他与上一行的非空常量值相同
	l       = len("Hello World") //常量的值也可以是编译器能够计算出结果的某些表达式的值
)

func main() {
	var x = 20
	const y = 20 //常量不一定需要使用
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T, %v \n", m, m)

	var a ANC = 30 //自定义类型
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
}
