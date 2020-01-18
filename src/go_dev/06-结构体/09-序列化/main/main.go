package main

import (
	"encoding/json"
	"fmt"
)

//
/*
结构体与 json
*/

type persoon struct {
	NameF string `json:"name_f"` // tag  使用 json 解析的时候使用的字段
	Age   int    `json:"age"`
}

func main() {
	// 结构体转 json   序列化
	p1 := persoon{
		"zhoulin",
		30,
	}

	res, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("struct marshal failed:%v", err.Error())
	}
	fmt.Printf("%v\n", string(res))

	// json 转为结构体   反序列化  并会把首字母大写

	str := `{"name_f":"leon","age":30}`
	var p2 persoon
	_ = json.Unmarshal([]byte(str), &p2) // struct 是值类型
	fmt.Printf("%#v\n", p2)

}
