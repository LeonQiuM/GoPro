package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a map[string]string
	// 不实例化、初始化直接操作会报错
	a = make(map[string]string)

	// 直接在声明时候初始化
	var b map[int]string = map[int]string{
		123: "aa",
		122: "bb",
	}

	// :=
	c := make(map[string]string)
	a["abc"] = "123"
	a["ab1"] = "122"
	c["aa"] = "aa"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fullMap()
	// map 是引用类型，

	// slice of map
	var x []map[string]string
	x = make([]map[string]string, 6)
	// 需要再次初始化map,未初始化为 nil
	if x[0] == nil {
		x[0] = make(map[string]string)
	}
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))

}

// 多个 map 嵌套
func fullMap() {
	fmap := make(map[string]map[string]string)
	fmap["persons"] = make(map[string]string)
	fmap["persons"]["name"] = "james"
	fmt.Println(fmap)
}
