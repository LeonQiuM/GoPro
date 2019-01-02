package main

import "fmt"

// for 循环

/*
几种形式
1. for 初始化;条件;表达式 {}
2. for 条件 {}
3. for {}   死循环

*/

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i, k := range arr {
		fmt.Println(i, k)
	}
}
