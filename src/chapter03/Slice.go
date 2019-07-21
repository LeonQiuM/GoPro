package main

import (
	"fmt"
)

//切片
/*
获取数组中的子数组，切片是不可比较的

*/
func main() {
	var numbers = []int{44, 3, 1, 67, 34}
	ni := numbers[1:3] // 跟源数组的值指向同一片内存空间,共享同一片内存空间
	fmt.Println(ni)
	ni[1] = 200
	fmt.Println(ni)
	fmt.Println(numbers) // 源数组的值也会被修改
	// 忽略开始索引和结束索引
	n2 := numbers[2:]
	fmt.Println(n2)
	n3 := numbers[:4] //  超过索引最大值直接回抛出异常
	fmt.Println(n3)
	// 迭代
	for _, n := range n3 {
		fmt.Println(n)
	}
	reverse(n3)
	fmt.Println(n3)      //  切片被翻转
	fmt.Println(numbers) // 源数组的切片位置也被翻转

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]

	}
}
