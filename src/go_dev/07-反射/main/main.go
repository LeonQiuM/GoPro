package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func main() {
	var a int = 200
	test(a)
	//var std Student = Student{
	//	Name:"jack",
	//}
	// test(std)
}

// 接口==类型<-->类

func test(x interface{}) {
	v := reflect.ValueOf(x) // 转换为reflect.Value类型后就可以对其进行分析
	// 传入的是 interface 接口类型，并不是实际的类型
	fmt.Println(reflect.TypeOf(x)) // 获取变量的类型
	fmt.Println("Kind", v.Kind())  // 常量

	/*
		对于一个 struct，TypeOf返回的是具体的那种类型，例如本例子的 Student 类型
		二对于 Student 类型，其实是一个结构体，是可以通过 Kind 获取的
	*/
	fmt.Println(reflect.ValueOf(x))                 // 获取变量的值
	fmt.Println(reflect.TypeOf(reflect.ValueOf(x))) // reflect.Value类型
	// fmt.Println(reflect.Value.Kind()) // 获取变量的类别，返回一个常量
	// reflect.Value.Interface() //转换为Interface类型
	// v 此时为reflect.Value类型，转为 interface 类型
	interfaceV := v.Interface()
	fmt.Println(reflect.TypeOf(interfaceV))

	// reflect.Value 类型获取对应的值
	fmt.Println("get value", v.Int())
	//fmt.Println("get value",v.Float())
	//fmt.Println("get value",v.String())
	//fmt.Println("get value",v.Type())
	//fmt.Println("get value",v.Bool())
	//fmt.Println("get value",v.Bytes())
}
