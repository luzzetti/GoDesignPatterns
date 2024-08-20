package main

import "fmt"

type BookType int

const (
	HardCover BookType = iota
	SoftCover
	PaperBack
	EBook
)

type Book struct {
	Name      string
	Author    string
	PageCount int
	Type      BookType
}

type Library struct {
	Collection []Book
}

// IterateBooks is a PUSH model iterator
func (l Library) IterateBooks(fn func(Book) error) {

	for _, b := range l.Collection {
		err := fn(b)
		if err != nil {
			fmt.Println("Error encountered", err)
			break
		}
	}
}

// createIterator follows a PULL model
func (l Library) createIterator() iterator {
	return &BookIterator{
		books: l.Collection,
	}
}

// Hardcoded sample data
var lib *Library = &Library{Collection: []Book{
	{
		Name:      "War and peace",
		Author:    "Leo Tolstoy",
		PageCount: 864,
		Type:      HardCover,
	},
	{
		Name:      "Crime and Punishment",
		Author:    "Leo Tolstoy",
		PageCount: 1225,
		Type:      SoftCover,
	},
	{
		Name:      "My little pony",
		Author:    "Unknown",
		PageCount: 10,
		Type:      EBook,
	},
}}
