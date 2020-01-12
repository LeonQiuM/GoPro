package main

import "fmt"

// Student docs
type Student struct {
	name  string
	age   int
	score int
}

// People docs
type People struct {
	float64 // 匿名字段
	Student // 匿名字段
}

func (student *Student) study(course string) {
	fmt.Printf("Student:%s studying %s\n", student.name, course)

}

func main() {
	var stu1 People
	stu1.name = "jack"
	stu1.float64 = 33.6
	fmt.Printf("we have %f mil people\n", stu1.float64)
	stu1.study("math")
}
