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
const maxChanSize = 50000

type FileLogger struct {
	LogLevel    LEVEL
	filePath    string
	fileName    string
	maxFileSize int64 // B
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *logMsg // 不存字符串，单个字符串占用内存太大
}

type logMsg struct {
	level     LEVEL
	Msg       string
	funcName  string
	line      int
	fileName  string
	timestamp string
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		fmt.Println(err)
	}
	fl := &FileLogger{
		LogLevel:    level,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, maxChanSize),
	}
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

func (f *FileLogger) writeLogBackground() {
	for {
		if f.checkSize(f.fileObj) {
			// 检查大小，是否需要切换日志
			newfile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newfile
		}

		select {
		case logTmp := <-f.logChan: // 把日志信息从chan中取出来，当娶不到日志的话也会阻塞
			logInfo := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n",
				logTmp.timestamp,
				UnParseLevel(logTmp.level),
				logTmp.fileName,
				logTmp.funcName,
				logTmp.line,
				logTmp.Msg)
			_, _ = fmt.Fprintf(f.fileObj, logInfo)
			if f.checkSize(f.errFileObj) {
				newfile, err := f.splitFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newfile
			}
			if logTmp.level >= ERROR {
				_, _ = fmt.Fprintf(f.errFileObj, logInfo)
			}
		default:
			// 从chan中取不到日志
			time.Sleep(50 * time.Millisecond)
		}
	}
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
	// 开启单独goroutine来用于写日志
	go f.writeLogBackground()
	return nil
}

func (l *FileLogger) enable(ll LEVEL) bool {
	return l.LogLevel <= ll
}
func (f *FileLogger) FilelogSave(lv LEVEL, msg string, arg ...interface{}) {
	if f.enable(lv) { // 需要被记录的日志级别
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		msgAll := fmt.Sprintf(msg, arg...)
		// 造一个msg的对象
		logTmp := &logMsg{
			level:     lv,
			Msg:       msgAll,
			funcName:  funcName,
			line:      lineNo,
			fileName:  fileName,
			timestamp: now.Format("2006_01_02_15:04:05.000"),
		}
		// 将日志写入到chan中，当chan满了后，会阻塞，所以使用select
		select {
		case f.logChan <- logTmp:
		default:
			// 什么也不做，保证业务代码不阻塞，但是极端情况会丢失日志
		}
	}
}

func (f *FileLogger) splitFile(fileObj *os.File) (*os.File, error) {
	// 需要切割：
	// 1. rename，
	nowStr := time.Now().Format("2006_01_02_15:04:05")
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
