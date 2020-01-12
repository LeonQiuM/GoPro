package main

import (
	"fmt"
	"io"
	"os"
)

// 文件打开读写例子

func main() {
	fileObj, err := os.Open("./file.txt")

	if err != nil {
		fmt.Println("open file failed,err: %v", err)
		return
	}
	defer fileObj.Close()
	var tmp = make([]byte, 128)
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
		}
		// fmt.Printf("current read %d bytes\n", n)
		fmt.Printf(string(tmp[:n]))
		if n < 128 {
			return
		}
	}

}
