package book

import "github.com/qbitty/design/pattern/iterator"

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bsItor *BookShelfIterator) HasNext() bool {
	return bsItor.index < bsItor.bookShelf.Length()
}

func (bsItor *BookShelfIterator) Next() (interface{}, error) {
	element, err := bsItor.bookShelf.BookAt(bsItor.index)
	if err != nil {
		return nil, err
	}

	bsItor.index++
	return element, nil
}

var _ iterator.Iterator = (*BookShelfIterator)(nil)
