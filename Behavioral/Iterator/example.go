package main

import "fmt"

func main() {

	//lib.IterateBooks(myBookCallback)

	// PUSH model
	lib.IterateBooks(func(book Book) error {
		fmt.Println("Book Title:", book.Name)
		return nil
	})

	// PULL model
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Println("Book Title:", book.Name)
	}

}

//func myBookCallback(b Book) error {
//	fmt.Println("Book Title:", b.Name)
//	return nil
//}
