package main

import (
	"encoding/json"
	"fmt"
	"tutorial/golang/test_structure/models"
)

// I : here you tell us what I is
type I interface {
}

func main() {
	var book1 = new(models.Books) /* Declare Book1 of type Book */
	var book2 = models.Books{}    /* Declare Book2 of type Book */
	var book3 = models.Books{Title: "Java Programming", Author: "Peer Talyor", Subject: "Java Programming Tutorial", BookID: 1234567}
	var book4 models.Books
	var book5 = book3
	var ptr *models.Books

	describe(book1)
	/* book 1 specification */
	(*book1).Title = "Go Programming"
	book1.Author = "Mahesh Kumar"
	book1.Subject = "Go Programming Tutorial"
	book1.BookID = 6495407

	describe(book2)
	/* book 2 specification */
	book2.Title = "Telecom Billing"
	book2.Author = "Zara Ali"
	book2.Subject = "Telecom Billing Tutorial"
	book2.BookID = 6495700

	/* print Book1 info */
	fmt.Printf("Print Book1 info\n")
	fmt.Printf("Book title : %s\n", book1.Title)
	fmt.Printf("Book author : %s\n", book1.Author)
	fmt.Printf("Book subject : %s\n", book1.Subject)
	fmt.Printf("Book book_id : %d\n", book1.BookID)

	/* print Book2 info */
	fmt.Printf("Print Book2 info\n")
	fmt.Printf("Book title : %s\n", book2.Title)
	fmt.Printf("Book author : %s\n", book2.Author)
	fmt.Printf("Book subject : %s\n", book2.Subject)
	fmt.Printf("Book book_id : %d\n", book2.BookID)

	describe(book3)
	fmt.Printf("Print Book3 info\n")
	book3.PrintBook()

	data, error := json.MarshalIndent(book3, "", "	")
	if error != nil {
		fmt.Printf("Encoding error: %v\n", error)
	}
	fmt.Printf("%s \n[%T]\n", data, data)

	error = json.Unmarshal(data, &book4)
	if error != nil {
		fmt.Printf("Decoding error: %v\n", error)
	}
	fmt.Printf("Print Book4 info (copy of Book3)\n")
	book4.PrintBook()

	describe(book4)
	fmt.Printf("Print Book5 info (copy of Book4)\n")
	ptr = &book4
	ptr.Subject = "Advance Java Programming Tutorial"
	(*ptr).BookID = 1234568
	(*ptr).PrintBook()

	book3.BookID = 1357924
	fmt.Printf("%v/%v\n", book3.BookID, book5.BookID)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
