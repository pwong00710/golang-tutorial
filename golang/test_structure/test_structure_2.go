package main

import (
	"tutorial/golang/test_structure/models"
)

func main() {
	var book1 models.Books /* Declare Book1 of type Book */
	var book2 models.Books /* Declare Book2 of type Book */

	/* book 1 specification */
	book1.Title = "Go Programming"
	book1.Author = "Mahesh Kumar"
	book1.Subject = "Go Programming Tutorial"
	book1.BookID = 6495407

	/* book 2 specification */
	book2.Title = "Telecom Billing"
	book2.Author = "Zara Ali"
	book2.Subject = "Telecom Billing Tutorial"
	book2.BookID = 6495700

	/* print Book1 info */
	book1.PrintBook()
	book2.PrintBook()

	/* print Book2 info */
}

/*
func printBook(book models.Books) {
	fmt.Printf("Book title : %s\n", book.Title)
	fmt.Printf("Book author : %s\n", book.Author)
	fmt.Printf("Book subject : %s\n", book.Subject)
	fmt.Printf("Book book_id : %d\n", book.BookID)
}
*/
