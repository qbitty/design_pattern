package aggregate

import (
	"github.com/qbitty/design/pattern/book"
	"github.com/qbitty/design/pattern/iterator"
)

type Aggregate interface {
	Iterator() iterator.Iterator
}

var _ Aggregate = (*book.BookShelf)(nil)
