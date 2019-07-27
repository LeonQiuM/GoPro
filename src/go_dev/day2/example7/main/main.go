package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	var a1 float32
	var b bool
	var c byte = 'a'
	a = 333
	fmt.Printf("%v\n", a)  //  %v  原样输出
	fmt.Printf("%v\n", c)  //  %v  原样输出
	fmt.Printf("%p\n", &a) //  %v  指针
	fmt.Printf("%T\n", b)  //  %v  类型输出
	fmt.Printf("%t\n", b)  //  %v  布尔型
	fmt.Printf("%b\n", a)  //  %v  二进制
	fmt.Printf("%d\n", a)  //  %v  十进制
	fmt.Printf("%o\n", a)  //  %v  八进制
	fmt.Printf("%f\n", a1) //  %v  浮点数

	/*
		%%%%%%
		q%  双引号围绕的字符串
		%x  16进制

	*/
	str := fmt.Sprintf("%d", a)
	fmt.Printf(str + "\n")
	fmt.Println(reflect.TypeOf(str))
	StrAdd()

}

//  字符串拼接

func StrAdd() {
	str1 := "Hello"
	str2 := "Jack"
	str3 := "World"

	fmt.Printf(str2 + str1 + str3)
	str := fmt.Sprintf("%s %s", str1, str3)
	fmt.Println(str)

	n := len(str1)
	fmt.Println(n)

	//  切片
	Hello := str[:5]
	fmt.Println(Hello)

	// 翻转
	fmt.Println(reverse("abcdefg"))
	fmt.Println(reverse1("abcdefg1111"))

}

func reverse(str string) string {
	var res string
	resLen := len(str)
	for i := 0; i < resLen; i++ {
		res = res + fmt.Sprintf("%c", str[resLen-i-1])
	}
	return res
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(tmp)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)

}
