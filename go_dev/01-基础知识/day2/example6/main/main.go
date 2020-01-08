package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //  初始化随机种子,当前时间的纳秒
	for i := 1; i < 11; i++ {
		roundNum := rand.Int()
		fmt.Println(roundNum)  
	}

	for i := 1; i < 11; i++ {
		roundNum := rand.Intn(100)
		fmt.Println(roundNum)
	}
	for i := 1; i < 11; i++ {
		roundNum := rand.Float64()
		fmt.Println(roundNum)
	}

}
