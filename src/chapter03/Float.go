package main

import "fmt"

// 浮点数

/*
1. 浮点数的表示方法
	种类：
		1. float32 6个10进制精度，小数点后可以精确到6位
				有效位整数部分为23位，其他的9位用于表示符号位，质数，
		2. float64 15个10进制精度，小数点后可以精确到15位
2. 格式化输出
3. 浮点数中的特殊值

*/
func main() {
	var f float32 = 1 << 24
	fmt.Println(f)
	// 省略
	const e = 2.718888
	const x1 = 0.12
	const x2 = .12
	const x3 = 43.00
	const x4 = 43.
	// 科学计数法
	const value1 = 1.23e21
	const value2 = 1.23E21
	// 格式化输出

	const v1 = 32323334.12322222223345555555555555
	fmt.Printf("科学计数法v=%g\r\n", v1)

}
