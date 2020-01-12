package MYLOGGER

import (
	"fmt"
	"strings"
	"time"
)

//
/*

 */
// Logger docs
type LEVEL uint16

const (
	// 定义日志级别
	TRACE LEVEL = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger struct {
	LogLevel LEVEL
}

func parseLogLevel(s string) (LEVEL, error) {
	s = strings.ToLower(s)

	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := fmt.Errorf("failed log level,use default info lovel")
		return INFO, err
	}
}

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

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s] [DEBUG] %s\n", now.Format("2006-01-02 15:04:05.000"), msg)
	}

}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		fmt.Printf("[%s] [INFO]  %s\n", now.Format("2006-01-02 15:04:05.000"), msg)
	}

}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] [WARNING]  %s\n", now.Format("2006-01-02 15:04:05.000"), msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s] [ERROR]  %s\n", now.Format("2006-01-02 15:04:05.000"), msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s] [FATAL]  %s\n", now.Format("2006-01-02 15:04:05.000"), msg)
	}
}
