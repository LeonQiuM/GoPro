package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(6) // 默认为cpu的逻辑核心数量，默认跑满整个cpu
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func a() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		time.Sleep(time.Millisecond * 30)
		fmt.Println("A:", i)
	}
}
func b() {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		time.Sleep(time.Millisecond * 30)
		fmt.Println("B:", i)
	}
}
