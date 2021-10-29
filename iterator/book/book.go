package book

type Book struct {
	name string
}

func (book *Book) GetName() string {
	return book.name
}
