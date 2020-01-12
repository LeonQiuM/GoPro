package MYLOGGER

import (
	"fmt"
	"os"
	"path"
	"time"
)

//
/*

 */
type FileLogger struct {
	LogLevel    LEVEL
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		fmt.Println(err)
	}
	fl := &FileLogger{}
	fl.LogLevel = level
	fl.filePath = fp
	fl.fileName = fn
	fl.maxFileSize = maxSize
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl

}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) initFile() error {
	fullName := path.Join(f.filePath, f.fileName)

	fileObj, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open logs file err,err:%v", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open err logs file err,err:%v", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (l FileLogger) enable(ll LEVEL) bool {
	return l.LogLevel <= ll
}
func (f *FileLogger) FilelogSave(lv LEVEL, msg string, arg ...interface{}) {
	if f.enable(lv) {

		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		msgAll := fmt.Sprintf(msg, arg...)
		_, _ = fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2006-01-02 15:04:05.000"),
			UnParseLevel(lv),
			fileName,
			funcName,
			lineNo,
			msgAll)
		if lv >= ERROR {
			_, _ = fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n",
				now.Format("2006-01-02 15:04:05.000"),
				UnParseLevel(lv),
				fileName,
				funcName,
				lineNo,
				msgAll)
		}
	}
}

func (f FileLogger) Debug(msg string, arg ...interface{}) {
	f.FilelogSave(DEBUG, msg, arg...)

}

func (f FileLogger) Info(msg string, arg ...interface{}) {
	f.FilelogSave(INFO, msg, arg...)

}

func (f FileLogger) Warning(msg string, arg ...interface{}) {
	f.FilelogSave(WARNING, msg, arg...)
}

func (f FileLogger) Error(msg string, arg ...interface{}) {
	f.FilelogSave(ERROR, msg, arg...)
}

func (f FileLogger) Fatal(msg string, arg ...interface{}) {
	f.FilelogSave(FATAL, msg, arg...)
}
