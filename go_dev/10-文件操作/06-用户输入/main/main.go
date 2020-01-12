package main

import (
	"bufio"
	"fmt"
	"os"
)

//
/*
获取用户输入是有空格
*/
func main() {
	useBufio()

}

func useScan() {
	var s string
	fmt.Printf("Please input:")
	n, err := fmt.Scanln(&s)
	fmt.Println(n, err)
	fmt.Printf("input:%s\n", s) // 无法输入带有空格的
}

func useBufio() { // bufio来读取带空格用户数输入
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	fmt.Printf("input:%s", s)

}
