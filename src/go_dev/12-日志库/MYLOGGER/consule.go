package MYLOGGER

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	LogLevel LEVEL
}

// NowLog docs
func NowConsoleLog(loglevel string) ConsoleLogger {
	level, err := parseLogLevel(loglevel)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		level,
	}

}

func (l ConsoleLogger) enable(ll LEVEL) bool {
	return l.LogLevel <= ll
}
func (l ConsoleLogger) logSave(lv LEVEL, msg string, arg ...interface{}) {
	if l.enable(lv) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		msgAll := fmt.Sprintf(msg, arg...)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2006-01-02 15:04:05.000"),
			UnParseLevel(lv),
			fileName,
			funcName,
			lineNo,
			msgAll)
	}
}

func (l ConsoleLogger) Debug(msg string, arg ...interface{}) {
	l.logSave(DEBUG, msg, arg...)

}

func (l ConsoleLogger) Info(msg string, arg ...interface{}) {
	l.logSave(INFO, msg, arg...)

}

func (l ConsoleLogger) Warning(msg string, arg ...interface{}) {
	l.logSave(WARNING, msg, arg...)
}

func (l ConsoleLogger) Error(msg string, arg ...interface{}) {
	l.logSave(ERROR, msg, arg...)
}

func (l ConsoleLogger) Fatal(msg string, arg ...interface{}) {
	l.logSave(FATAL, msg, arg...)
}
