package main

import (
	"errors"
	"fmt"
)

//  函数多返回值
/*

 */
func calc(m, n int) (result, mm, nn int) {
	result = m + n
	mm = m
	nn = n
	return

}

func MyErr(m int) (str string, err error) {
	if m >= 0 {
		str = "Hello World"
		err = nil

	} else {
		str = string(m)
		err = errors.New("m 不能小于 0")

	}
	return

}

func main() {
	result, m, n := calc(12, 20)
	fmt.Println("result:", result, "m:", m, "n", n)

	res, err := MyErr(-3)
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
