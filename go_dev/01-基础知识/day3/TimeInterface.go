package main

import (
	"fmt"
	"time"
)

<<<<<<< HEAD
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
=======
func main() {
	//  Date 输出指定时区的指定时间
	t := time.Date(2019, time.November, 21, 10, 10, 33, 1, time.UTC)
	fmt.Printf("%s\n", t.Local())
	fmt.Println(t.Location()) //   时区

	// Now 当地时间对象
	t1 := time.Now()
	fmt.Printf("%s\n", t1)
>>>>>>> 4dbc1c723d7158468317847d692dbebece639d59
}
