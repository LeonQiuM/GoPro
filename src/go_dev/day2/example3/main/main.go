package main

import (
	"fmt"
	"time"
)

const (
	Female = 2
	Man    = 1
)

func main() {
	for {
		second := time.Now().Unix()
		fmt.Println(second)
		if second%Female == 0 {
			fmt.Printf("female\n")
		} else {
			fmt.Printf("man\n")
		}
		time.Sleep(1000 * time.Millisecond)

	}

}
