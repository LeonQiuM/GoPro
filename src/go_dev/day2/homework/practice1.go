package main

import "fmt"

//  101-200 之间的所有素数，并输出所有的素数

func main() {
	for i := 101; i <= 200; i++ {
		var flag bool = true // 默认是素数
		for j := 2; j < i; j++ {
			if i%j == 0 { //  除以 2到i-1 之间的每一个，只要结果等于 0，就说明有余数
				flag = false // 就把状态修改为不是素数，直接跳出
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}

}
