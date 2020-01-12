package MYLOGGER

import "fmt"

//
/*

 */
// Logger docs
type Logger struct {
}

// NowLog docs
func NowLog() Logger {
	return Logger{}

}

func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	fmt.Println(msg)
}

func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}
