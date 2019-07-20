package main

import "fmt"

//常量
/*
常量表达式的值是在编译器计算的，而不是在运行时，为了减少运行时的错误，也为了编译优化
常量不适用不会出错
常量类型必须是基础类型，
bool、string、数字
*/
func main() {
	const pi = 3.1415926
	//  批量声明多个常量
	const (
		e    = 2.1111
		mark = "flag"
	)

	// 变量的批量声明
	var (
		x int
		y = 230.1
	)
	fmt.Println(x, y)

	// 常量的所有算术运算、逻辑运算、比较运算、类型转换的结果也是常量

	const arrayLen = 20

	var names [arrayLen]string // 声明一个 string类型的 20 长度的数组
	fmt.Println(names)
	/*
		var arrayLen1 = 33
		var names1 [arrayLen1]string    // 数组的声明也是在编译器声明，数组的长度不能使用变量，需要使用常量，数组的长度一般不变
		fmt.Println(names1)
	*/

}
