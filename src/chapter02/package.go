package main

import (
	"fmt"
	"product"
)

// 包

func main() {
	fmt.Println("aaa")
	//fmt就是一个包
	// fmt 在 src 里面有一个目录叫做 fmt
	// 默认的所有路径就是 GOROOT
	// 自定义的就去 GOPATH 中搜索，GOPATH 可以设定全局（Global PATH）的和项目的（Project PAth）
	// /Users/qiumeng/GoglandProjects/GoPro/library  引入到 src 的上一级目录即可
	fmt.Println(product.GerNmme())
	fmt.Println(product.Nmme)
}
