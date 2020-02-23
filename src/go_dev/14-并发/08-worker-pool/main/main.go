package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	for k := range jobs {
		//fmt.Printf("worker:%d start job:%d\n", id, k)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, k)
		results <- k * 2 // 将结果写入channel
	}
	wg.Done()
}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	// 多个任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
	}()
	// 启动一个goroutine去一直取
	go func() {
		for x := range result { // 通道不关闭，for-range会一直取，这样就会造成死锁
			fmt.Println(x)
		}
		fmt.Println("result close, exit")
	}()

	// 开启三个goroutine，去执行所有的任务
	wg.Add(3)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}
	wg.Wait()
	fmt.Println("all jobs done,close results")
	close(result)
	/*
		for a := 1; a <= 5; a++ {
			res := <-result
			fmt.Println(res)
		}
	*/
}
