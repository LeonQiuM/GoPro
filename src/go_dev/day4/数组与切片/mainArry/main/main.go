package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(a[0])
	// fmt.Println(a[10])  //  越界

	for index, value := range a {
		fmt.Printf("index:%d  value%d \n", index, value)
	}

	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	arr2[0] = 100
	fmt.Println(arr1)
	fmt.Println(arr2)

	// 在其他函数中修改数组，需要传入地址
	var x [5]int
	changeValue(&x)
	fmt.Println(x)

}

func changeValue(arr *[5]int) {
	(*arr)[0] = 1000
}
