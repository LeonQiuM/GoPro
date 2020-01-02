package main

import "fmt"

//
/*
接口的嵌套

*/
type ReadWriter interface {
	Read()
	Write()
}

type Look interface {
	Look()
	UnLook()
}

type File struct {
	ReadWriter
	Look
}

func (f *File) Read() {
	fmt.Println("read file")

}

func (f *File) Write() {
	fmt.Println("Write file")

}

func main() {
	var f File
	f.Read()
	f.Write()

}
