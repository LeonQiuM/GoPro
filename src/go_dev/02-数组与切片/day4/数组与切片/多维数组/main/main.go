package main

import "fmt"

func main() {
	testArray2()
}

func testArray2() {
	var a [2][5]int = [...][5]int{
		{1, 2, 3, 4, 5},
		{23, 543, 13, 1, 2},
	}
	//  多维数组的遍历
	for index, v := range a {
		for k, i := range v {
			fmt.Printf("%d,%d=%d ", index, k, i)
		}
		fmt.Println()
	}
}
