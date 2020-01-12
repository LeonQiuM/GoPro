package main

import (
	"fmt"
	"strconv"
)

// 自动以数据类型
/*
目的： 用更有意义的数据类型名称保存变量的值

*/
type Year int16
type Month int8
type Day int8
type Salay float32
type Flag bool

func printInfo(year Year, month Month, day Day) {
	fmt.Println(
		strconv.Itoa(int(year)) + "年" + strconv.Itoa(int(month)) + "月" + strconv.Itoa(int(day)) + "日")
}

func main() {
	var year Year = 2018
	var month Month = 12
	var day Day = 33
	printInfo(year, month, day)

	var salary Salay = 333.3333
	var flag Flag = true
	if flag {
		fmt.Println(salary)
	}
}
