package main

import "fmt"

// 有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？分别是多少？

func main() {
	var count int
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			for k := 1; k <= 4; k++ {
				if (i != j) && (i != k) && (k != j) {
					count += 1
					fmt.Println(i, j, k)
				}
			}
		}
	}
	fmt.Println(count)
}
