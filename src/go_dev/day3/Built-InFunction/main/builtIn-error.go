package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("Init config failed")
}

func test() {
	/*
			defer func() { // 使用 recover 将异常进行捕获
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()


		b := 0
		a := 100 / b //  会直接吧异常抛出
		fmt.Println(a)
	*/

	//  自己抛出异常
	err := initConfig()
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	for {
		test()
		time.Sleep(time.Second * 2)
	}

	fmt.Println("111111")

}
