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
	maxFileSize int64 // B
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

func (l *FileLogger) enable(ll LEVEL) bool {
	return l.LogLevel <= ll
}
func (f *FileLogger) FilelogSave(lv LEVEL, msg string, arg ...interface{}) {
	if f.enable(lv) {

		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		msgAll := fmt.Sprintf(msg, arg...)
		if f.checkSize(f.fileObj) {
			newfile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newfile
		}
		_, _ = fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2006-01-02 15:04:05.000"),
			UnParseLevel(lv),
			fileName,
			funcName,
			lineNo,
			msgAll)
		if f.checkSize(f.errFileObj) {
			newfile, err := f.splitFile(f.errFileObj)
			if err != nil {
				return
			}
			f.errFileObj = newfile
		}
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

func (f *FileLogger) splitFile(fileObj *os.File) (*os.File, error) {
	// 需要切割：
	// 1. rename，
	nowStr := time.Now().Format("2006_01_02_15:04:05.000")
	fullName := path.Join(f.filePath, fileObj.Name())
	newLogName := fmt.Sprintf("%s_%s", fullName, nowStr)
	_ = os.Rename(fullName, newLogName)
	//2. 关闭当前，
	fileObj.Close()
	// 3. 新建一个文件对象
	newFileObj, err := os.OpenFile(fullName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new file err:%v\n", err)
		return nil, err
	}
	return newFileObj, nil
}

func (f *FileLogger) checkSize(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		return false
	}
	return info.Size() > f.maxFileSize // 如果当前文件大小大于最大设置值，需要切割，返回 true

}

func (f *FileLogger) Debug(msg string, arg ...interface{}) {
	f.FilelogSave(DEBUG, msg, arg...)

}

func (f *FileLogger) Info(msg string, arg ...interface{}) {
	f.FilelogSave(INFO, msg, arg...)

}

func (f *FileLogger) Warning(msg string, arg ...interface{}) {
	f.FilelogSave(WARNING, msg, arg...)
}

func (f *FileLogger) Error(msg string, arg ...interface{}) {
	f.FilelogSave(ERROR, msg, arg...)
}

func (f *FileLogger) Fatal(msg string, arg ...interface{}) {
	f.FilelogSave(FATAL, msg, arg...)
}
