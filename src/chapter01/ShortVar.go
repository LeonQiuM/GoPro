package main

import (
	"fmt"
)

// 简短变量

/*
1.必须在函数内申明
2.var 申明变量，可以不初始化
3. 简短变量必须初始化
4. 不需要提供数据类型
*/

var x = 100

func main() {
	fmt.Println(&x, x) // 全局变量
	x := "111"         // 重新定义初始化同名局部变量
	y := 0.24
	fmt.Printf("x=%s,y=%f\n", x, y)
	fmt.Println(&x)

	a, b, c := 1, 0.2, "abc" //  局部变量
	fmt.Println(a, b, c)
	roolabck_same_zone()
	roolabck_diff_zone()
}

//  简短变量退化赋值，简短变量退化赋值前提是必须在同一个作用域内

func roolabck_same_zone() {
	x := 100
	fmt.Println(&x, x)
	x, y := 200, "edf" //  此时退化为赋值，由于x的类型已经被定义，此处只可以修改为相同类型的值
	fmt.Println(&x, x) // 内存地址相同可以确定x为同一变量
	fmt.Println(y)
}

func roolabck_diff_zone() {
	/*
		退化赋值在不同的作用域，两个都是全新的变量
	*/
	x := 100
	fmt.Println(&x, x)
	{
		x := 200
		fmt.Println(&x, x)
	}
}
