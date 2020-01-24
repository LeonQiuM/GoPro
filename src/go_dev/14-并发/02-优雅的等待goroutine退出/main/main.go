package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup // 是一个结构体，结构体是值类型，不能复制

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	wg.Wait()
}

func f() {
	for i := 0; i < 5; i++ {
		rand.Seed(time.Now().UnixNano()) // 保证每次执行都使用一个随机数种子
		r1 := rand.Int()
		r2 := rand.Intn(100)
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	fmt.Println(i)
}

/*
goroutine 什么时候结束
	+ goroutine对应的函数结束返回，这个goroutine就结束了
*/
