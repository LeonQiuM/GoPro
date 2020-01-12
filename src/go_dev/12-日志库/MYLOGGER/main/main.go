package main

// 自定义一个日志库

import (
	"errors"
	"go_dev/12-日志库/MYLOGGER"
	"time"
)

func main() {

	log := MYLOGGER.NewFileLogger("info", "./", "app.log", 10*1024)

	for {
		log.Debug("debug log")
		log.Info("info log")
		log.Warning("warning log")
		log.Error("Error log,err:%s", errors.New("err get user name"))
		log.Fatal("Fatal log")
		time.Sleep(time.Second * 2)
	}

}
