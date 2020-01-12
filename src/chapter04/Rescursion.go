package main

import "fmt"

//函数递归
/*

 */

func jc(n uint) uint {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * jc(n-1)
	}
}
func main() {
	fmt.Println(jc(3))
}
