package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
bufio读取
*/
func main() {
	fileObj, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("open file failed,err: %v", err)
		return
	}
	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			fmt.Println("read line err")
			return
		} else {
			fmt.Printf(line)
		}
	}
}
