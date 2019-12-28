package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// 字符换转int，有两个返回值，有err
	num, err := strconv.Atoi("1231231")
	if err == nil {
		fmt.Printf("num:%d,type:%s\n", num, reflect.TypeOf(num))
	} else {
		fmt.Println("FATAL change type err,", err)
	}

	//  int 转string，  只有一个返回值，无err
	num1 := strconv.Itoa(1231)
	fmt.Printf("num:%s,type:%s\n", num1, reflect.TypeOf(num1))

	// ParseInt 将字符串转换为对应类型的函数
	num2, err2 := strconv.ParseInt("23341", 10, 64)
	/*
		参数1 数字的字符串形式
		参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
		参数3 返回结果的bit大小 也就是int8 int16 int32 int64
	*/
	if err2 == nil {
		fmt.Printf("num:%d,type:%s\n", num2, reflect.TypeOf(num2))
	} else {
		fmt.Println("FATAL change type err,", err2)
	}

	// ParseFloat  将字符串转换为对应类型的函数
	num3, err3 := strconv.ParseFloat("2.323341", 64)
	if err3 == nil {
		fmt.Printf("num:%f,type:%s\n", num3, reflect.TypeOf(num3))
	} else {
		fmt.Println("FATAL change type err,", err3)
	}
	// ParseUint  将字符串转换为对应类型的函数
	num4, err4 := strconv.ParseUint("2", 10, 64)
	if err4 == nil {
		fmt.Printf("num:%d,type:%s\n", num4, reflect.TypeOf(num4))
	} else {
		fmt.Println("FATAL change type err,", err4)
	}
	//  ParseBool 将字符串转换为对应类型的函数
	num5, err5 := strconv.ParseBool("true")
	if err5 == nil {
		fmt.Printf("num:%v,type:%s\n", num5, reflect.TypeOf(num5))
	} else {
		fmt.Println("FATAL change type err,", err5)
	}

	// 将整数转换为字符串  func FormatInt(i int64, base int) string
	fmt.Printf("%s", strconv.FormatInt(234545, 10))

}
