package MYLOGGER

import (
	"fmt"
	"path"
	"runtime"
	"strings"
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

func UnParseLevel(l LEVEL) string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	}
	return "INFO"
}

func getInfo(skip int) (funcName, fileName string, LineNumber int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.CAller() failed,err:%v\n", ok)

	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
	fileName = path.Base(file)
	LineNumber = line
	return

}
