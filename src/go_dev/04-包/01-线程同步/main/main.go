package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var lock sync.Mutex

func main() {

	var a map[int]int
	a = make(map[int]int)
	a[0] = 11
	a[1] = 22
	a[2] = 33

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[0] = rand.Intn(100)
			lock.Unlock()

		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()
}
