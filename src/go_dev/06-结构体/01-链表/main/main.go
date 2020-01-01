package main

import (
	"fmt"
)

// Student docs
type Student struct {
	Name  string
	Age   int
	Score int
	Next  *Student
}

func main() {
	var head Student
	head.Age = 18
	head.Name = "jack"
	head.Score = 100

	var stu1 Student
	stu1.Age = 18
	stu1.Name = "jack1"
	stu1.Score = 100

	head.Next = &stu1

	var p *Student = &head
	for p != nil {
		fmt.Println(*p)
		p = p.Next

	}

	var stu2 Student
	stu2.Age = 18
	stu2.Name = "jack2"
	stu2.Score = 100

	stu1.Next = &stu2 // 链表插入
}
