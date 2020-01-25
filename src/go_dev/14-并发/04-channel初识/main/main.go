package main

import (
	"fmt"
	"sync"
)

/*单纯的函数并发没有意义
函数与函数之间需要做数据件韩才能体现并发执行函数的意义
+ 共享内存：性能问题
+ goroutine是并发执行的执行体，channel是他们之间的连接

channel是一种特殊的类型
*/
var wg sync.WaitGroup

func main() {
	// 定义一个channel
	var c chan int     // int类型的channel，需要制定同道中元素的类型。channel是引用类型
	fmt.Println(c)     // nil
	c = make(chan int) // 通道的初始化后才能使用，引用类型使用make初始化【slice、map、channel】
	wg.Add(1)
	go func() {
		x := <-c
		fmt.Println("get x:", x)
		wg.Done()
	}()
	c <- 10 // 放一个10进去
	fmt.Println("10 put to channel")
	// c = make(chan int, 16) // 带缓冲区的初始化
	fmt.Println(c)
	wg.Wait()
	// 发送值
	// 接收值
	// 关闭
	close(c)
}
