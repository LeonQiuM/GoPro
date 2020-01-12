package main

// 自定义一个日志库

import (
	"go_dev/12-日志库/MYLOGGER"
	"time"
)

func main() {

	log := MYLOGGER.NowLog()

	for {
		log.Debug("debug log")
		log.Info("info log")
		log.Warning("warning log")
		time.Sleep(time.Second * 2)
	}

}
