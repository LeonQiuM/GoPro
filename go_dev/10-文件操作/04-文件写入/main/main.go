package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/*
+ os.O_WRONLY  只写
+ os.O_CREATE   创建文件
+ os.O_RDONLY		只读
+ os.O_RDWR		读写
+ os.O_TRUNC		 清空
+ os.O_APPEND		追加
*/

// perm 文件权限，0777、0755等设定
func main() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	fileObj.Write([]byte("jakc like milk\n"))
	fileObj.WriteString("jake sh xx.sh")
	fileObj.Close()

	// xxx
	bufioWrite()
	// xxxx
	writeFile()
}

func bufioWrite() {
	file, err := os.OpenFile("xxx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file err:%v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("Golang\n")
	}
	writer.Flush() // 将缓存中的内容写入文件，用户态到内核态
}

func writeFile() {
	str := "hello Beijing"
	// 文件不存在会创建
	err := ioutil.WriteFile("xxxx.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf(" write file err: %v", err)
		return
	}
}
