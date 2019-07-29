package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() { 
	//  int 转字符串
	fmt.Println("int --> str:"+strconv.Itoa(1111111), reflect.TypeOf(strconv.Itoa(1111111)))

	// 字符换转int
	num, err := strconv.Atoi("1231231")
	if err == nil {
		fmt.Println("num:%s", num, reflect.TypeOf(num))
	} else {
		fmt.Println("FATAL change type err,", err)
	}

}
