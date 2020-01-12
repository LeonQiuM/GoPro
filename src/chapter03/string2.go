package main

import (
	"fmt"
	"unicode/utf8"
)

//unicode 和 utf8 编码
/*
ucs-2
ucs-4

*/
func main() {
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c") // tuf8
	fmt.Println("\u4e16\u754c")             // ucs-2
	fmt.Println("\U00004e16\U0000754c")     // ucs-4

	s := "hello 张三"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s)) //  获得字符数

}
