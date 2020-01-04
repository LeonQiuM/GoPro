package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"student_name"`
	Age  int    `json:"student_age"`
}

func main() {
	var a int = 200
	test(a)
	//var std Student = Student{
	//	Name:"jack",
	//}
	// test(std)

	xxx := 123
	setTest(&xxx)
	// 通过反射设置,但是无法改变值类型的副本的值，所以最好传入指针

	// 通过反射操作结构体
	var as Student = Student{
		Name: "Alex",
	}
	structTest(as)
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

func setTest(xx interface{}) {
	value := reflect.ValueOf(xx)
	// reflect.Value 通过反射来设置变量的值
	value.Elem().SetInt(100) // 通过反射设置,但是无法改变值类型的副本的值，所以最好传入指针
	// reflect 中无法使用指针，故提供了 Elem方法来操作指针指向的变量
	fmt.Println(value.Elem())
}

func (student Student) Print() {
	fmt.Println(student)
}

//func (student Student) Create(name string) {
//	student.Name = name
//}

func structTest(as interface{}) {
	va := reflect.ValueOf(as)
	kd := va.Kind()
	if kd != reflect.Struct {
		// 常量类型不是结构体
		fmt.Println("expect struct")
		return
	}
	fieldsNum := va.NumField() // 有多少个字段
	fmt.Println("fieldsNum", fieldsNum)
	methodNum := va.NumMethod() // 有多少中方法
	fmt.Println("methodNum", methodNum)
	// 通过反射来调用结构体中的方法
	var reflectVauleSlice []reflect.Value
	va.Method(0).Call(reflectVauleSlice)
	// 通过反射获取多个字段的值
	for i := 0; i < fieldsNum; i++ {
		fmt.Printf("index[%d],value[%v],type[%v]\n", i, va.Field(i), va.Field(i).Kind())
	}
	// 修改 struct 的字段名称,修改的话传入的需要是一个指针
	// va.Elem().Field(0).SetString("James")
	// 获取 struct 的 tag 相关,tag 是类型的一部分
	studentNameTag := reflect.TypeOf(as).Field(1).Tag.Get("json")
	fmt.Println(studentNameTag)

}
