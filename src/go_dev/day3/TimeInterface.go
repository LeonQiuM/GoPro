package main

import (
	"fmt"
	"time"
)

func main() {
	//  Date 输出指定时区的指定时间
	t := time.Date(2019, time.November, 21, 10, 10, 33, 1, time.UTC)
	fmt.Printf("%s\n", t.Local())
	fmt.Println(t.Location()) //   时区

	// Now 当地时间对象
	t1 := time.Now()
	fmt.Printf("%s\n", t1)
}
