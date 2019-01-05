package main

import "fmt"

// bool类型
/*

 */

func main() {
	var b bool = true
	var a bool = false
	fmt.Println(a, b)
	if b {
		fmt.Println(a)
	}
	s := "hello"
	if s != "" && s[1] == 'e' {
		fmt.Println("OK")
	}
}
