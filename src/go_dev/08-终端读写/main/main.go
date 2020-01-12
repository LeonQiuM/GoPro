package main

import (
	"fmt"
	"os"
	"time"
)

//
/*

 */
func main() {
	for {
		time.Sleep(time.Second)
		file, err := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("not such file or dir")
		}
		defer file.Close()
		//fmt.Fprintf(os.Stdout, "xxx\n")
		fmt.Fprint(file, "xxxx\n")
	}
}
