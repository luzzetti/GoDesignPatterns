package main

type IterableCollection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	next() *Book
}

// BookIterator structure keeps track of where my iterator is
type BookIterator struct {
	current int
	books   []Book
}

func (i *BookIterator) hasNext() bool {
	if i.current < len(i.books) {
		return true
	}
	return false
}

func (i *BookIterator) next() *Book {
	if i.hasNext() {
		bk := i.books[i.current]
		i.current++
		return &bk
	}
	return nil
}
