package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwlock sync.RWMutex

func main() {

	var count int32
	var a map[int]int
	a = make(map[int]int)
	a[0] = 11
	a[1] = 22
	a[2] = 33

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwlock.Lock()
			b[0] = rand.Intn(100)
			rwlock.Unlock()

		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			rwlock.RLock()
			fmt.Println(b)
			rwlock.RUnlock()
			count++
			// atomic.AddInt32(&count, 1)  // 原子操作
		}(a)
	}
	time.Sleep(time.Second * 3)
	//fmt.Println(atomic.LoadInt32(&count))  // 原子操作
	fmt.Println(count)
}
