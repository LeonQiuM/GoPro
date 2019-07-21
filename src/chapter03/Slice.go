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

	// 手动比较切片
	s := []string{"abc", "xyz", "bbb", "abc", "xyz"}
	s1 := s[0:2]
	s2 := s[3:]
	//  fmt.Println(s1==s2)  // 不能直接比较
	fmt.Println(s1, s2)
	fmt.Println(equal(s1, s2))

	// 但是可以判断切片是否为空
	fmt.Println(s1 == nil)

	//  使用 make 开创建切片
	ok1 := make([]int, 20) //  并且各个初始化为 0 值
	ok2 := make([]string, 20)
	fmt.Println(ok1, ok2)

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]

	}
}

func equal(x, y []string) bool {
	// 两个分片的比较
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true

}
