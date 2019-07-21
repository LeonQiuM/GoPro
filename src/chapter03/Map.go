package main

import (
	"fmt"
	"reflect"
	"sort"
)

//映射 map
/*
是一种 hash 表结构，是一中 key-value 结构，每个 key 都不相同，是无序的
在 map 中检索数据的时间复杂度是一个常数
*/
func main() {
	ages := map[string]int{ // [key_类型]value——类型
		"Mike": 33,
		"John": 23,
	}
	fmt.Println(ages)
	ages1 := make(map[string]int) //  空的 hash 表
	ages1["Mike"] = 22
	ages1["Jack"] = 33
	fmt.Println(ages1)

	// 检索
	fmt.Println(ages1["Mike"])
	//  检索一个不存在的 key
	fmt.Println(ages1["abc"]) //  不存在会返回 0

	// 删除
	delete(ages1, "Mike")
	fmt.Println(ages1)
	delete(ages1, "Mike1") // 删除的值不存在，不会报错
	fmt.Println(ages1)
	// fmt.Println(&ages1["Mike"])  // 不能取地址
	// 迭代
	/*
		names := []string{"a", "b"}
		for i, v := range names {
			fmt.Println(i, v) //  索引和值
		}
	*/

	ages["cc"] = 1
	ages["cc1"] = 12
	ages["cc"] = 12 //  值的修改
	for name, age := range ages {
		fmt.Println(name, age) //  map 无序，故意的，不想让程序果粉依赖 hash 函数的实现
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println(names)
	sort.Strings(names) // 排序
	fmt.Println(names)

	for _, name := range names {
		fmt.Println(name, ages[name]) //  排序后再按照顺序输出
	}

	//  map 类型的零值，表示不指向类型的 hash 表
	lingMap := make(map[string]int)
	fmt.Println(lingMap)
	lingMap = nil
	fmt.Println(lingMap == nil, reflect.TypeOf(lingMap), len(lingMap))
	// 查找、删除、len、range 都可以安全的工作在 nil 之上
	//  不能为一个空值map设置值
	// lingMap["ssss"]=11  // panic: assignment to entry in nil map

	// 区分 value 为 0 还是值本身就是 0
	ages["xy"] = 0
	age, ok := ages["aa"]
	age1, ok1 := ages["xy"]
	fmt.Println(age, ok)
	fmt.Println(age1, ok1)
	if age, ok := ages["xyz"]; !ok {
		fmt.Println(age, "不存在")
	}

	// map 的比较，不能直接进行比较
	fmt.Println(ages, ages1)
	mapEqual(ages, ages1)
}

func mapEqual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if yv, ok := y[k]; !ok || yv != v {
			return false
		}
	}
	return true
}
