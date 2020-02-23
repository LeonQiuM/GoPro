package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel

func main() {
	ch := make(chan int, 1)
	//ch <- 100
	//<-ch
	//
	//fmt.Println(x, ok)
	go Send(ch) // 用另外一个goroutine启动这个函数，每5s往chan中放一个值
	for {
		x, ok := <-ch
		fmt.Println(x, ok) // 当chan被关闭的时候，取的值是类型的零值，ok为false，更好的方式为for-range
		time.Sleep(time.Second)  // 本goroutine，每秒去取一次，会等待4s才能取到一个其他goroutine生成的数
	}
}

func Send(ch chan<- int) {
	for {
		num := rand.Intn(1000)
		ch <- num
		time.Sleep(5 * time.Second)
	}
}
