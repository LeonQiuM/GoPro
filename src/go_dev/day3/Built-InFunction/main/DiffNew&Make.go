package main

import "fmt"

func test1() {
	s1 := new([]int)     // 使用 new 申请
	*s1 = make([]int, 5) // 使用 make 进行初始化
	(*s1)[0] = 100
	fmt.Println(s1)

	s2 := make([]int, 10)
	s2[0] = 200
	fmt.Println(s2)
	return

}

func main() {
	test1()

}
