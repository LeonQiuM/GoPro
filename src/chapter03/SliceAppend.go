package main

import "fmt"

// 向切片中追加元素,append 函数
/*

 */
func main() {
	var x = []int{1, 2} // 切片
	var y = []int{1, 2} // 切片
	// var y = []int{1,2}  // 切片 ，不定义长度即为切片，定义了长度为数组
	x1 := x[0:2]
	fmt.Println(x1)
	fmt.Println(&x[1])
	xx := append(x1, 3, 5, 6, 7, 8, 9) // 重新创建一个分片,所以内存指向不相同
	fmt.Println(x1)
	fmt.Println(xx)
	fmt.Println(&xx[1])

	// 切片不能直接相加,也不能使用 append 方法直接把一个列表添加到另外一个的末尾
	for _, v := range y {
		x = append(x, v)
	}
	fmt.Println(x)

}
