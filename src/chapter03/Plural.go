package main

import (
	"fmt"
	"unsafe"
)

// 复数

/*
10 + 5i
实部+虚部
complex64 相当于实部是 float32、虚部也是 float32
complex128  相当于实部是 float64、虚部也是 float64

函数的两个参数的类型必须统一
*/

func main() {
	var r1 float32 = 20
	var i1 float32 = 12
	var x complex64 = complex(r1, i1)
	fmt.Println(x) //(20+12i)
	fmt.Println(unsafe.Sizeof(x), "字节")
	var r2 float64 = 34
	var i2 float64 = 13
	var x2 complex128 = complex(r2, i2)
	fmt.Println(x2) //(20+12i)
	fmt.Println(unsafe.Sizeof(x2), "字节")

	// 复数的运算
	var x3 complex128 = complex(3, 5)
	var x4 complex128 = complex(2, 4)
	fmt.Println(x3 + x4) // (5+9i)
	fmt.Println(x3 - x4) // (1+1i)
	fmt.Println(x3 * x4) // (-14+22i)
	/*复数的乘法
	z1 = a+bi
	z2 = c+di
	z1*z2 = (ac-bd)+(bc+ad)i
	*/
	fmt.Println(x3 / x4) // (1.3-0.1i)
	/*复数的乘法
	z1 = a+bi
	z2 = c+di
	z1/z2 = (ac+bd)/(c^2+d^2)+((bd-ad)/(c^2+d^2))
	*/
	fmt.Println(real(x3 * x4)) // 直接复数的实部和虚部
	fmt.Println(imag(x3 * x4)) // 直接复数的实部和虚部

	x11 := 3 + 12i
	y := 12 + 33i
	fmt.Println(x11)
	fmt.Println(y)

}
