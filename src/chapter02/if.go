package main

import "fmt"

// if 条件语句
/*
if
else if
else if
else
*/
func main() {
	i := -1
	flag := true
	if i > 5 || flag {
		fmt.Println("OK")
	} else if i < 0 {
		fmt.Println("test")
	} else {
		fmt.Println("end")
	}
}
