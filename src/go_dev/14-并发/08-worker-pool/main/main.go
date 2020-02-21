package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for k := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, k)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, k)
		results <- k * 2 // 将结果写入channel
	}
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	// 开启三个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		res := <-result
		fmt.Println(res)
	}

}
