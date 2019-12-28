package main

import "fmt"

/*
对于一个数字 n，球 n 的阶乘之和
*/

func main() {
	var num int
	fmt.Scanf("%d", &num)
	var sum int
	var tmp int = 1
	for i := 1; i <= num; i++ {
		tmp = i * tmp
		sum = tmp + sum
	}
	fmt.Println(sum)

}
