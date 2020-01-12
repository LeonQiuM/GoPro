package main

import (
	"fmt"
	"sort"
)

var a map[int]string

func main() {
	findInMap()
	rangeMap()
	fmt.Println("------")
	sortMap()
	fmt.Println("------")
	reverseMap()
}

// 插入和跟新

// 查找

func findInMap() {
	a = make(map[int]string, 20)
	a[1] = "name"
	a[2] = "age"
	a[3] = "hight"
	a[4] = "skin"
	fmt.Println(a)
	val, ok := a[5]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("not found")

	}

}

// 遍历

func rangeMap() {
	a = make(map[int]string, 20)
	a[1] = "name"
	a[2] = "age"
	a[3] = "hight"
	a[4] = "skin"
	for k, v := range a {
		fmt.Println(k, v)
	}
	delete(a, 1) // 根据某个 key 删除
	fmt.Println(len(a))
}

// 删除

//  长度

// map排序，需要自己支持，key 为无序
func sortMap() {
	var intmap map[int]int
	intmap = make(map[int]int)
	intmap[1] = 675
	intmap[2123] = 31257
	intmap[21235] = 31245
	intmap[1231] = 312314

	var keys []int
	for k := range intmap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, value := range keys {
		fmt.Println(value, intmap[value])
	}

}

// 反转
func reverseMap() {
	var intmap map[string]int
	var reversedMap map[int]string

	intmap = make(map[string]int)
	reversedMap = make(map[int]string)

	intmap["1"] = 675
	intmap["1231"] = 312314

	for k, v := range intmap {
		reversedMap[v] = k
	}
	fmt.Println(reversedMap)

}
