package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	context, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("read file failed,err:%v", err)
		return
	}
	fmt.Printf(string(context))
}
