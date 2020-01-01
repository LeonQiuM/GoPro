package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 200
	test(a)
}

func test(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	fmt.Println(reflect.TypeOf(reflect.ValueOf(x)))  // reflect.Value类型
}
