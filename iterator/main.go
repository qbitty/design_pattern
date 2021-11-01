package main

import (
	"fmt"
	"log"

	"github.com/qbitty/design/pattern/book"
)

func main() {
	bookshelf := book.New(3)
	bookshelf.Append(book.NewBook("C"))
	bookshelf.Append(book.NewBook("go"))
	bookshelf.Append(book.NewBook("kotlin"))
	iterator := bookshelf.Iterator()
	for iterator.HasNext() {
		ele, err := iterator.Next()
		if err != nil {
			log.Fatal(err)
		}
		if bk, ok := ele.(*book.Book); ok {
			fmt.Println(bk.GetName())
		}
	}
}
