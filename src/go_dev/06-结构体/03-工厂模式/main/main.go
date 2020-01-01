package model

// Student docs
type Student struct {
	Name string
	Age  int
}

// NewStudent comment
func NewStudent(name string, age int) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

/*
package main
S:=new(Student)
S:= NewStudent("Jack",22)
*/
