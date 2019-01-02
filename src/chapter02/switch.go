package main

import "fmt"

// switch 条件语句

/*
1. 每个 case 执行完成自动 break

*/

func main() {
	value := "Hello"
	switch value { // 可以加小括号，可以不加
	case "Hello":
		fmt.Println("你好")
		fallthrough // 想要让本次执行完成后继续向下执行，fallthrough不能用在左后一个 case 上
	case "world":
		fmt.Println("世界")
	case "what":
		fmt.Println("啥")
	}

	n := 20
	switch {
	case n < 10:
		fmt.Println("10")
	case n < 19:
		fmt.Println("19")
	case n >= 20:
		fmt.Println("29")
	default:
		fmt.Println("other")

	}
}
