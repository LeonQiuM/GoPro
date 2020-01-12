package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fileObj, err := os.OpenFile("./xxx.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Printf("open file failed,err:%vn", err)
	}

	log.SetOutput(fileObj)
	for {
		log.Println("test log~")
		time.Sleep(time.Second * 2)
	}
}
