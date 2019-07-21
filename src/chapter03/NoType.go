package main

import (
	"fmt"
	"math"
	"reflect"
)

//无类型常量和变量，根据表达式自动推算
/*

 */
func main() {
	var x float32 = math.Pi
	fmt.Println(reflect.TypeOf(x), x)
	var y float32 = math.Pi
	fmt.Println(reflect.TypeOf(y), y)

	//  不指定类型
	var p = math.Pi
	fmt.Println(reflect.TypeOf(p))

	var s = "Hello"
	fmt.Println(reflect.TypeOf(s))
}
