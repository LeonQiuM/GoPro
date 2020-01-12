package main

import "fmt"

type student struct {
	name string
}

//
/*
类型的断言，由于接口是一般类型，不知道其具体的类型，如果需要转成集体的类型，就可以采用类型的断言
*/
func main() {
	var a interface{}
	var b int
	a = b
	c, ok := a.(int) // 带检查的 interface 断言
	if ok {
		fmt.Println(c)
	}

}

// 参数类型的断言
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) { // type  为关键字
		case bool:
			fmt.Printf("param #%d is bool\n", i)
		case float64, float32:
			fmt.Printf("param #%d is float64\n", i)
		case int:
			fmt.Printf("param #%d is int\n", i)
		case string:
			fmt.Printf("param #%d is string\n", i)
		case student: // 结构体的断言
			fmt.Printf("param #%d is Student\n", i)
		case *student: // 结构体指针的断言
			fmt.Printf("param #%d is *Student\n", i)
		default:
			fmt.Printf("UnKown type")

		}
	}
}
