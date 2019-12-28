package main

import (
	"fmt"
)

func main() {
	initSlice()

	// 切片无法进行比较
	// 数组可以进行直接比较
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	if a1 == a2 {
	}

	findSlicePtr()
}

func findSlicePtr() {
	// 切片的指针是指向数组的第一个元素
	var x = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := x[1:5] // 一个切片
	// 打印的地址是完全仙童的
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &x[1])

}

func initSlice() {
	// 定义和初始化
	var testSlice []int

	// 从数组定义切片
	var arr [5]int = [5]int{0, 1, 2, 3, 4}
	testSlice = arr[1:5]
	fmt.Println(testSlice)
	//  len 计算长度
	//  cap计算容量
	fmt.Println(len(testSlice))
	fmt.Println(cap(testSlice))

	testSlice = arr[0:1]
	fmt.Println(len(testSlice)) // 1
	fmt.Println(cap(testSlice)) // 5

}
