package main

import (
	"fmt"
	"go_dev/06-结构体/07-接口/04-loadBalance/balance"
	"os"
	"time"
)

//
/*

 */

func main() {
	insts := make([]balance.Instance, 0)
	for i := 0; i < 20; i++ {
		one := balance.Instance{
			Host: fmt.Sprintf("10.111.19.%d", i+2),
			Port: 8000 + i,
		}
		insts = append(insts, one)
	}

	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}
	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Printf("do balance err, get inst failed:%s\n", err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Printf("Post data to http://%s:%d\n", inst.Host, inst.Port)
		// fmt.Println(inst) // 自己实现了struct 的 String 方法
		time.Sleep(time.Second)
	}

}
