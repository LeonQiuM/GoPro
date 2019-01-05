package main

import (
	"fmt"
	"math"
)

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
	fmt.Printf("科学计数法：v=%g\r\n", v1)
	fmt.Printf("科学计数法精确到6位：v=%e\r\n", v1)
	fmt.Printf("6位：v=%f\r\n", v1)
	fmt.Printf("输出前面留位，不够补空格，并保留一定的精度：v=%20.7f\r\n", v1)
	fmt.Printf("保留一定的精度：v=%.7f\r\n", v1)
	// 特殊值

	var ok float64 // 默认为0.0
	fmt.Println(ok, -ok, 1/ok, -1/ok, ok/ok)
	fmt.Printf("1/ok=+Inf(正无穷):%t\r\n", math.IsInf(1/ok, 1))
	fmt.Printf("-1/ok=+Inf(负无穷):%t\r\n", math.IsInf(-1/ok, -1))
	fmt.Printf("ok/ok=+Inf(NaN):%t\r\n", math.IsNaN(ok/ok))
	nan := math.NaN()
	inf := math.Inf(-1)
	fmt.Println(nan, inf)
}
