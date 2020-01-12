package MYLOGGER

import (
	"fmt"
	"time"
)

// NowLog docs
func NowLog(loglevel string) Logger {
	level, err := parseLogLevel(loglevel)
	if err != nil {
		panic(err)
	}
	return Logger{
		level,
	}

}

func (l Logger) enable(ll LEVEL) bool {
	return l.LogLevel <= ll
}
func logSave(lv LEVEL, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",
		now.Format("2006-01-02 15:04:05.000"),
		UnParseLevel(lv),
		fileName,
		funcName,
		lineNo,
		msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		logSave(DEBUG, msg)
	}

}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		logSave(INFO, msg)
	}

}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		logSave(WARNING, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		logSave(ERROR, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		logSave(FATAL, msg)
	}
}
