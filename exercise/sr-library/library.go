//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	Title string
	CheckedOut string
	CheckeOutTime time.Time
	ReturneTime time.Time
}

type Member struct {
	Name string
	RentedBooks []Book
}

type BookStore struct {
	Members []Member
	Books []Book
}

func LogInfo(bookStore *BookStore) {
	fmt.Println()
	fmt.Println("Member info:")
	for _, member := range bookStore.Members {
		fmt.Printf("\t%s's rented books:\n", member.Name)
		for _, book := range member.RentedBooks {
			fmt.Printf("\t\t%s:\n\t\tCheckout time: %s\n", book.Title, book.CheckeOutTime)
		}
	}
	fmt.Println()
	fmt.Println("Book info:")
	for _, book := range bookStore.Books {
		fmt.Printf("%s\n\tCheckout: %s\n\tReturn: %s\n", book.Title, book.CheckeOutTime, book.ReturneTime)
		if book.CheckedOut != "" {
			fmt.Printf("\tChecked out by: %s\n", book.CheckedOut)
		}
	}
}

func CheckOutBook(book *Book, member *Member) {
	if !book.CheckeOutTime.IsZero() {
		fmt.Println("Book already checked out!")
	}

	book.CheckeOutTime = time.Now()
	book.CheckedOut = member.Name
	book.ReturneTime = time.Time{}
	member.RentedBooks = append(member.RentedBooks, *book)
}

func CheckInBook(book *Book, member *Member) {
	if book.CheckeOutTime.IsZero() {
		fmt.Println("Book already returned!")
	}

	book.ReturneTime = time.Now()
	book.CheckedOut = ""
	book.CheckeOutTime = time.Time{}

	newList := []Book{}
	for _, rentedBook := range member.RentedBooks {
		if rentedBook.Title != book.Title {
			newList = append(newList, rentedBook)
		} 
	}

	member.RentedBooks = newList
}

func main() {
	bookStore := BookStore{
		Members: []Member{
			{
				Name: "Pijus",
			},
			{
				Name: "Algirdas",
			},
			{
				Name: "Matas",
			},
		},
		Books: []Book{
			{
				Title: "Alice in wonderland",
			},
			{
				Title: "Cars",
			},
			{
				Title: "Harry potter",
			},
			{
				Title: "Rich dad, poor dad",
			},
		},
	}

	CheckOutBook(&bookStore.Books[2], &bookStore.Members[0])
	LogInfo(&bookStore)
	CheckOutBook(&bookStore.Books[1], &bookStore.Members[1])
	LogInfo(&bookStore)
	CheckInBook(&bookStore.Books[2], &bookStore.Members[0])
	LogInfo(&bookStore)
}
