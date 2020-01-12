package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "Hello World abc\n"
	fmt.Printf(strings.Replace(str, "World", "you", 1))
	fmt.Printf("%d", strings.Count(str, "World"))
	fmt.Printf(strings.Repeat(str, 3))
	fmt.Printf(strings.ToLower(str))
	fmt.Printf(strings.ToUpper(str))
	fmt.Printf(strings.TrimSpace(str)) // 去掉两头的空

}
