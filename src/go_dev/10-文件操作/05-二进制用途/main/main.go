package main

import (
	"fmt"
)

// 111
/*
吃饭  		100
睡觉		010
打豆豆 		 001
*/

const (
	eat   int = 4
	sleep int = 2
	play  int = 1
)

func main() {
	f(eat | play)         // 101
	f(eat | play | sleep) // 101
}
func f(arg int) {
	fmt.Printf("%b\n", arg)
}
