package main

import (
	"fmt"
	"reflect"
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

	//
	var Slice1 []int
	var array1 [3]int
	fmt.Println(reflect.TypeOf(Slice1))
	fmt.Println(reflect.TypeOf(array1))
	fmt.Printf("%p\n", &Slice1)
	fmt.Printf("%p\n", &array1)

	// 两个切片的 append
	var x1 = []int{1, 2, 3}
	var x2 = []int{4, 5, 6, 7}
	x1 = append(x1, x2...)
	fmt.Println(x1)

	// 切片的 copy,copy 的容量为设定的容量，超过会丢弃掉，大了会是默认值
	var a = []int{1, 2, 3, 4, 5, 6}
	var b = make([]int, 10)
	copy(b, a)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	// string 与 slice ，string 的底层就是一个 byte 的数组，因此可以进行切片操作
	s := "hello world"
	s1 := s[0:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)

	// 字符串中的字符本身并不能够被修改，但是将其转化为数组，再进行修改
	sx := "hello world"
	sx1 := []rune(sx) // 字节数量可变，用于支持中文，byte 是只能在英文下使用的
	sx1[1] = '8'
	sxChangede := string(sx1)
	fmt.Println(sxChangede)

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
