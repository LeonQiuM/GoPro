package main

import (
	"fmt"
)

func main() {
	recusive(0)
	x := factor(4)
	fmt.Println(x)

}

func factor(n int) int {
	if n == 1 {
		return 1
	}
	return factor(n-1) * n
}

func recusive(n int) {
	fmt.Println("Hello")
	// time.Sleep(time.Second * 1)
	if n > 10 {
		return
	}
	recusive(n + 1)

}
