package models

import "fmt"

// Books : here you tell us what Book is
type Books struct {
	Title   string
	Author  string
	Subject string
	BookID  int
}

// PrintBook : here you tell us what PrintBook is
func (book Books) PrintBook() {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book book_id : %d\n", book.BookID)
}
