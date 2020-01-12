package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	getInfo(0)

}

func getInfo(n int) {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Printf("runtime.CAller() failed,err:%v\n", ok)

	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}
