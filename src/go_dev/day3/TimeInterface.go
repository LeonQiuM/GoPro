package main

import (
	"fmt"
	"time"
)

func S() {
	time.Sleep(time.Millisecond * 200)
}

func main() {
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02 3:04"))
	fmt.Println(now.Format("2006/1/02 3:04"))

	start := time.Now().UnixNano()
	S()
	end := time.Now().UnixNano()
	fmt.Printf("COST[%d ms]\n", (end-start)/1000000)
}
