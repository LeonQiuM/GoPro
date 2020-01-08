package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaaa")
	//var a int
	a := 10
	fmt.Println(a)
	b_a := &a
	fmt.Println(b_a) // b_a 为指向a的指针
	*b_a = 100
	fmt.Println(a)
}
