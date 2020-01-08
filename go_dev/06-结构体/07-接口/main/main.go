package main

import "fmt"

type Test interface {
	// 接口定义，定于 struct 的一组方法

	// 定义的接口必须要全部实现
	// 空接口可以保存任何类型的类型
	echoName()
}

type Hello interface {
	// 如果一个变量还有了多个 interface 类型的方法，那么这个变量就实现了多个接口
	Hello()
}

type Student struct {
	name string
	age  int
}

func (student *Student) echoName() { // 需要申明是哪个 struct 的方法
	fmt.Println(student.name)
}

func (student *Student) Hello() {
	fmt.Printf("hello, my name is %s\n", student.name)
}

func main() {
	var stu = Student{
		name: "jack",
		age:  32,
	}
	var t Test
	t = &stu
	t.echoName()
	var h Hello
	h = &stu
	h.Hello()

}
