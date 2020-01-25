package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 启动三个goroutine
	for w := 1; w <= 3; w++ {
		go work(w, jobs, results)
	}
	// 有5个任务
	for j := 0; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//获取结果
	for a := 0; a <= 5; a++ {
		resp := <-results
		fmt.Println(resp)
	}
}

func work(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second) // do
		fmt.Printf("worker:%d done job:%d\n", id, j)
		// 吧结果返回到通道
		result <- j * j

	}
}
