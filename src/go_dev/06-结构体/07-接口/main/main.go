package main

import "fmt"

type Test interface {
	// 接口定义，定于 struct 的一组方法

	// 定义的接口必须要全部实现
	echoName()
}

type Student struct {
	name string
	age  int
}

func (student *Student) echoName() {
	fmt.Println(student.name)
}

func main() {
	var stu = Student{
		name: "jack",
		age:  32,
	}
	var t Test
	t = &stu
	t.echoName()

}
