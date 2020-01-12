package main

// 自定义一个日志库

import (
	"go_dev/12-日志库/MYLOGGER"
	"time"
)

func main() {

	log := MYLOGGER.NowLog("info")

	for {
		log.Debug("debug log")
		log.Info("info log")
		log.Warning("warning log")
		log.Error("Error log")
		log.Fatal("Fatal log")
		time.Sleep(time.Second * 2)
	}

}
