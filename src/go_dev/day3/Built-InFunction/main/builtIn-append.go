package main

import "fmt"

func main() {
	var a []int
	a = append(a, 10, 20, 30)
	fmt.Println(a)
	a = append(a, a...) // 两个 slince 拼接起来
	fmt.Println(a)
}
