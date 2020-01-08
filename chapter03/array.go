package main

import (
	"fmt"
	"reflect"
)

//数组
/*
1. 数组的声明
2. 枚举
3. 数组的比较
声明后的数组长度不可变

var arrayName[length] type
*/
func main() {
	var a [10]int
	fmt.Println(a)    //  不为数组初始化，那么其每一个值都是响应类型的 0 值
	fmt.Println(a[1]) //  第二个值
	fmt.Println(a[len(a)-1])
	for i := range a {
		a[i] = i * i
	}
	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	fmt.Println(a)

	// 声明时初始化
	var a1 [3]int = [3]int{4, 5, 6}
	var a2 [3]int = [3]int{1, 3} // 只初始化一部分
	fmt.Println(a1)
	fmt.Println(a2)

	var a3 = [3]int{1, 2, 3}
	fmt.Println(a3)
	//
	var a4 = [...]int{4, 11, 23, 22, 3, 5, 1, 2, 3, 55, 123, 45234} // 使用三个...  来不指定长度
	fmt.Println(a4)
	// 数组的长度必须是常量，不能是变量

	// key-value声明
	symbol := [...]string{10: "abc"} // 第 10 个(最后一个) index 的值为 "abc",前面的都为空串
	symbol1 := [...]string{10: "abc", 20: "xyz", 30: "123"}
	fmt.Println(reflect.TypeOf(symbol), symbol)
	fmt.Println(reflect.TypeOf(symbol1), symbol1)

	var changeValue = [...]int{1, 2, 3, 5}
	changeValue[0] = 3

	fmt.Println(changeValue)

	//  数组的比较
	arr1 := [2]int{1, 2}
	arr2 := [...]int{1, 2}
	fmt.Println(arr1 == arr2)

	//
	arr := [...]int{1, 2, 3, 5, 4}
	fmt.Println(fun(arr))
	fmt.Println(arr) //  函数内部修改不影响源

	fun1(&arr) // 通过将源的函数指针传入，可以在新的函数中修改值
	fmt.Println(arr)
}

func fun(ptr [5]int) int { // 复制一份压入函数栈
	var result int
	ptr[1] = 333
	for _, v := range ptr {
		result += v
	}
	return result
}

func fun1(ptr *[5]int) {
	for i, v := range ptr {
		ptr[i] = v * v

	}
}
