package books

import "fmt"

var bookList = []map[string]string{}

func AddBook(title, author string) {
	book := map[string]string{"title": title, "author": author}
	bookList = append(bookList, book)
}

func ListBooks() {
	if len(bookList) == 0 {
		fmt.Println("No books available.")
		return
	}
	for i, book := range bookList {
		fmt.Printf("%d. %s by %s\n", i+1, book["title"], book["author"])
	}
}
