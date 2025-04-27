package main

import "fmt"

type Book struct {
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books []Book
}

func (e *Library) AddBook(Title string, Author string) {
	if Title != "" && Author != "" {
		addBook := Book{Title: Title, Author: Author, IsBorrowed: false}
		// add book to Lib :
		e.Books = append(e.Books, addBook)
		fmt.Printf("Book \"%v\" is added . Author \"%v\".\n", Title, Author)
	} else {
		fmt.Println("Author or Title can't be empty !")
	}

}

func (e *Library) BorrowBook(title string) bool {
	for item := range e.Books {
		if e.Books[item].Title == title {
			e.Books[item].IsBorrowed = true
			return true
		}
	}
	return false
}

func (e *Library) ReturnBook(title string) bool {
	for item := range e.Books {
		if e.Books[item].Title == title {
			e.Books[item].IsBorrowed = false
			return true
		}
	}
	return false
}

func (e *Library) ListAvailableBooks() []string {
	var array []string
	for item := range e.Books {
		if e.Books[item].IsBorrowed != true {
			array = append(array, e.Books[item].Title)
		}
	}
	return array
}

func main() {
	var Library Library

	Library.AddBook("Don Quixote", "Miguel de Cervantes")
	Library.AddBook("Alice's Adventures in Wonderland", "Lewis Carroll")
	Library.AddBook("The Adventures of Huckleberry Finn", "Mark Twain")
	Library.AddBook("The Adventures of Tom Sawyer", "Mark Twain")
	Library.AddBook("Treasure Island", "Robert Louis Stevenson")

	fmt.Println(Library.BorrowBook("Don Quixote"))
	fmt.Println(Library.BorrowBook("The Adventures of Tom Sawyer"))
	fmt.Println(Library.BorrowBook("Treasure Island"))

	fmt.Println(Library.ReturnBook("Don Quixote"))

	fmt.Println(Library.ListAvailableBooks())

}
