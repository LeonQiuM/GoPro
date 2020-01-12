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
	fLook()
	fUnLook()
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
	var f *File
	f.Read()
	f.Write()
	var test interface{}
	test = f
	v, ok := test.(ReadWriter) //判断一个结构体是否存在某个接口
	fmt.Println(ok, v)

}
