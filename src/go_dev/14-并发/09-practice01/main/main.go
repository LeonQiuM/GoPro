package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//
//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主goroutine从resultChan取出结果并打印到终端输出
//为了保证业务代码的执行性能将之前写的日志库改写为异步记录日志方式。

type Job struct {
	value int64
}

type Result struct {
	job *Job
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)
var wg sync.WaitGroup

func gen(ch1 chan<- *Job) {
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		ch1 <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func Consumer(ch1 <-chan *Job, rc chan<- *Result) {
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newRes := &Result{
			job: job,
			sum: sum,
		}
		resultChan <- newRes
	}

}

func main() {
	wg.Add(1)
	go gen(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go Consumer(jobChan, resultChan)
		wg.Done()
	}
	for result := range resultChan {
		fmt.Printf("Value:%d,sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
