package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		// 满足的case执行，多个满足是随机
		case x := <-ch:
			fmt.Printf("%d\n", x)
		case ch <- i:
		}
	}
}
