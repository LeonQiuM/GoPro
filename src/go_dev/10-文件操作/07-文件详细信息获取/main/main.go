package main

import (
	"fmt"
	"os"
)

// 文件的详细信息获取
func main() {

	file, _ := os.Open("./main.go")
	fmt.Printf("%T\n", file)

	info, _ := file.Stat()
	fmt.Printf("size:%d B\n", info.Size())

}
