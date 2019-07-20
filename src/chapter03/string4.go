package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//字符串与数字之间转换
/*
strconv
*/
func main() {
	// 整数转换为字符串
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(reflect.TypeOf(y), y)

	z := strconv.Itoa(x)
	fmt.Println(reflect.TypeOf(z), z)

	//  进制转换
	bin := strconv.FormatInt(int64(x), 2) // 十进制转换为 2 进制
	fmt.Println(reflect.TypeOf(bin), bin)

	oct := strconv.FormatInt(int64(x), 8) // 十进制转换为 8 进制
	fmt.Println(reflect.TypeOf(oct), oct)

	hex := strconv.FormatInt(int64(x), 16) // 十进制转换为 8 进制
	fmt.Println(reflect.TypeOf(hex), hex)

	_t := strconv.FormatInt(int64(x), 3) // 十进制转换为 3 进制
	fmt.Println(reflect.TypeOf(_t), _t)

	//  way2
	bin1 := fmt.Sprintf("%b", x)
	oct1 := fmt.Sprintf("%o", x)
	hex1 := fmt.Sprintf("%x", x)
	fmt.Println(reflect.TypeOf(bin1), bin1)
	fmt.Println(reflect.TypeOf(oct1), oct1)
	fmt.Println(reflect.TypeOf(hex1), hex1)

	//  字符串转换为数字
	s := "a456"
	x1, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println(x1)
	} else {
		fmt.Println("error")
	}

	// 16进制转为 10 进制
	x3, err := strconv.ParseInt("ABCD", 16, 64)
	if err == nil {
		fmt.Println(x3)
		fmt.Println(reflect.TypeOf(x3))
	} else {
		fmt.Println("error")
	}

	var a, c int
	var b string
	fmt.Scanf("%d,%x,%b", &a, &b, &c)
	fmt.Println(a, b, c)

}
