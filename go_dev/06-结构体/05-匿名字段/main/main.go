package main

import (
	"fmt"
	"time"
)

// Car 匿名字段
/*
结构体中的字段可以没有名字，即匿名字段
*/
type Car struct {
	Name string
}

// Train docs
type Train struct {
	Car
	start time.Time
	int   // 类型匿名
}

func main() {
	var train Train
	train.Name = "huoche"
	train.int = 200
	fmt.Println(train.int)
}

// 无冲突可以省略
