package main

import "fmt"

//  结构体
/*
是一种聚合的数据类型，是有 0 个或者任意多个各种类型的变量聚合而成的实体每个变量称为结构体的成员

经典案例:员工信息

*/
func main() {
	type Employee struct {
		ID       int
		Name     string
		address  string
		Position string
		Salary   float64
		email    [2]string
	}
	var bill Employee
	bill.ID = 10
	bill.Name = "Bill"
	bill.address = "BeiJing"
	bill.Position = "aaa"
	bill.Salary = 111111.1111
	bill.email[0] = "bill@abc.com"

	//  使用指针修改
	fmt.Println(bill)
	position := &bill.Position
	*position = "Shanghai"
	fmt.Println(bill)

	var employee *Employee = &bill
	employee.Position = "地球"
	fmt.Println(employee)
}
