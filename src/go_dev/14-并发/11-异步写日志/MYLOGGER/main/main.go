package main

// 自定义一个日志库

import (
	"GoPro/src/go_dev/14-并发/11-异步写日志/MYLOGGER"
	"errors"
)

var logger MYLOGGER.Logger

func main() {

	//log := MYLOGGER.NewFileLogger("info", "./", "app.log", 10*1024)
	logger = MYLOGGER.NewFileLogger("info", "./", "app.log", 10*1024*1024)
	// logger = MYLOGGER.NowConsoleLog("info")

	for {
		logger.Debug("debug log")
		logger.Info("info log")
		logger.Warning("warning log")
		logger.Error("Error log,err:%s", errors.New("err get user name"))
		logger.Fatal("Fatal log")
		// time.Sleep(time.Second * 1)
	}

}
