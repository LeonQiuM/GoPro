package main

import (
	"fmt"
	"sync"
)

var one sync.Once

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	go f1(a)
	go f2(a, b) // 两个消费者
	go f2(a, b)
	for ret := range b {
		fmt.Println(ret)
	}

	// 关闭channel
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1) // 关闭后可以取值，
	for x := range ch1 {
		fmt.Println(x)
	}
	// 超过界限后得到两个值,一个是对应类型的0值，一个是bool值，无值为false
	res, err := <-ch1
	fmt.Println(res, err)
}

/*
1. goroutine 1 生成100个数
2. goroutine 2 冲channel1中取值，计算平方，当发哦channel2中
3. 在main中从channel2中去除最终结果
*/

func f1(ch1 chan<- int) { // 定义一个单向通道，只可以向通道内发送值，一般在参数的时候时候使用
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) { // <-chan 只能取值
	//	chan<- 只写
	//	<-chan 只读
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	one.Do(func() {
		close(ch2)
	})
}
