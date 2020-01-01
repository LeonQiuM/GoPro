package main

import "fmt"

// Student docs
type Student struct {
	name  string
	age   int
	score int
}

func (student *Student) study(course string) {
	fmt.Printf("Student:%s studying %s\n", student.name, course)

}

func main() {
	var stu1 Student
	stu1.name = "jack"
	stu1.study("math")
}
