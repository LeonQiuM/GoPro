package model

type Student struct {
	Name  string
	Grade string
	Id    string
	Sex   string
	books []*Book
}
