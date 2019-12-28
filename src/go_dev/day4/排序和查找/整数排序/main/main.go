package main

import (
	"fmt"
	"sort"
)

func main() {
	// intSorted()
	// strSorted()
	// floatSorted()

	// 数组查找，查找必须是排序后的
	intSearched()

}

func intSorted() {
	var a = [...]int{1, 3, 254, 6, 786, 324, 234, 765678, 234}
	sort.Ints(a[:]) // 由于数组是值类型，传入的是数组的引用，所以传入的不能是数组，故需要传入切片
	fmt.Println(a)

}
func intSearched() {
	var a = [...]int{1, 3, 254, 6, 786, 324, 234, 765678, 239}
	index := sort.SearchInts(a[:], 234) // 查找返回的位置为排序后的位置
	sort.Ints(a[:])
	fmt.Println(a)
	index2 := sort.SearchInts(a[:], 234)
	fmt.Println(index)
	fmt.Println(index2)

}

func strSorted() {
	var str = [...]string{"abc", "1", "abz", "ann"}
	sort.Strings(str[:])
	fmt.Println(str)
}

func floatSorted() {
	var f = [...]float64{3.2, 11.22, 3.5435, 123.22}
	sort.Float64s(f[:])
	fmt.Println(f)
}
