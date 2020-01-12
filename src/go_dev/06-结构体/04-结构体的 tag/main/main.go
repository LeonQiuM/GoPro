package main

import (
	"encoding/json"
	"fmt"
)

// tag
type student struct {
	Name  string `json:"NameField"`
	Age   int    `json:"ageField"`
	Score int
	// 其他包无法访问本 package 中的小写字段
	// 但是如果非要给返回小写的，可以在 tag 中设定
}

/*
tag 可以通过反射获取到，最常用的场景为 json 的序列化和反序列化
*/

func main() {
	var stu = student{
		Name:  "jack",
		Age:   32,
		Score: 100,
	}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode err")
	}
	fmt.Println(string(data))

}
