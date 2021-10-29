package book

import (
	"errors"

	"github.com/qbitty/design/pattern/iterator"
)

type BookShelf struct {
	books []Book
	last  int
}

func (bs *BookShelf) Append(book Book) {
	bs.books[bs.last] = book
	bs.last++
}

func (bs *BookShelf) BookAt(index int) (*Book, error) {
	if bs.last <= 0 {
		return nil, errors.New("book shelf is empty")
	}

	if index > bs.last || index < 0 {
		return nil, errors.New("index our of bound")
	}

	return &bs.books[index], nil
}

func (bs *BookShelf) Length() int {
	return bs.last
}

func (bs *BookShelf) Iterator() iterator.Iterator {
	return &BookShelfIterator{
		bookShelf: bs,
	}
}
