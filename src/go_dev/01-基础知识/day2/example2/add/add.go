package add

import (
	"fmt"
	_ "go_dev/day2/example2/sub" // _ 只是引入这个包进行初始化,import 是最先被使用的
)

//  1. 全局变量最先初始化并复制
var Name string = "xxxx"
var Age int = 99
var InitName string

// 2. inti 函数执行，变量背覆盖
func init() {
	InitName = "Jack"
	Name = "Hello world"
	Age = 10
	fmt.Printf("init_name:%s,addr:%p\n", Name, &Name)
	fmt.Printf("init_Age:%d,addr:%p\n", Age, &Age)
}
