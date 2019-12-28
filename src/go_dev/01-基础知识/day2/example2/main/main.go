package main

import (
	"fmt"
	self_add "go_dev/day2/example2/add"
)

func main() {
	// init  被自动执行
	fmt.Printf("initName:%s\n", self_add.InitName)
	fmt.Printf("name:%s,addr:%p\n", self_add.Name, &self_add.Name)
	fmt.Printf("age:%d,%p\n", self_add.Age, &self_add.Age)
}
