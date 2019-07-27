package main

import (
	"fmt"
	"strconv"
)

/*
打印 100-999 之间的水仙花数 153=1*1*1+5*5*5+3*3*3
*/

func main() {
	//fmt.Println(23 / 10)
	for a := 100; a < 1000; a++ {
		i := a / 100 % 10
		j := a / 10 % 10
		k := a % 10
		// fmt.Println(i, j, k)
		sum := i*i*i + j*j*j + k*k*k
		// fmt.Println(sum)

		if sum == a {
			fmt.Println(a)
		}

	}

	fmt.Println("----------")
	flower()

}

func flower() {
	/*
		str := "abcdefghijk"
		for i := 0; i < len(str); i++ {
			fmt.Printf("%d\n", str[i])
		}
	*/
	var str string
	fmt.Scanf("%s", &str)
	var result int
	for i := 0; i < len(str); i++ {
		num := int(str[i] - '0') // 8的 ascii 为 38，  0 的 ascii为 30
		result += num * num * num

	}
	number, err := strconv.Atoi(str)
	if result == number && err == nil {
		fmt.Println("ok")
	}

}
