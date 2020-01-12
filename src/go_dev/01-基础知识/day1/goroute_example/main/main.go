package main

import (
	"fmt"
	"go_dev/day1/goroute_example/gorouter"
)

func main() {
	var pipe chan int
	pipe = make(chan int, 2)
	go gorouter.Add(100, 220, pipe)
	sum := <-pipe

	fmt.Printf("sum=%d\n", sum)

}
