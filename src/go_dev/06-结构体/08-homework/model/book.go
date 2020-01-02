package model

type Book struct {
	Name       string
	Total      int
	Author     string
	CreateTime string
}

var (
	ErrStockNotEnough = errors.new("not ecnought")
)

func CreateBook(name string, total int, author string, createTime string) (b *Book) {
	b = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateTime: createTime,
	}
	return b
}

func (b *Book) Canborrow(c int) bool {
	return b.Total > c
}

func (b *Book) Borrow(c int) error {
	if b.Canborrow(c) == false {
		err := ErrStockNotEnough
		return
	}
	b.Total -= c
	return
}

func (b *Book) RollBack(c int) error {
	b.Total += c
	return
}
