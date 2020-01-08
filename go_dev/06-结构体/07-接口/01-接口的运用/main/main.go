package main

import (
	"fmt"
	"sort"
)

//
/*
接口的运用
*/

type Student struct {
	Name string
	Age  int
	Id   string
}

type Book struct {
	Name   string
	Author string
}

type studentArray []Student

func (p studentArray) Len() int {
	// 返回长度
	return len(p)
}

func (p studentArray) Less(i, j int) bool {
	// 按照年龄排序
	return p[i].Age > p[j].Age
}

func (p studentArray) Swap(i, j int) {
	// 位置交换
	p[i], p[j] = p[j], p[i]
}

func main() {
	var stus studentArray
	for i := 0; i < 10; i++ {
		stu := Student{
			Name: fmt.Sprintf("stud%d", i),
			Age:  i + 120,
			Id:   string(i) + "1asdw3rt4g5r7u",
		}
		stus = append(stus, stu)
	}
	fmt.Println(stus)
	sort.Sort(stus) // sort 的接受参数为 interface，但是这个 interface 需要实现 Len、Less、Swap
	// 参考https://golang.org/pkg/sort/#Interface，排序的方式为 less 中实现的逻辑
	fmt.Println(stus)
}
