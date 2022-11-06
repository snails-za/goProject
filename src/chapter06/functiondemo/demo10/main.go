package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func changeBook(book Books) {
    book.title = "book1_change"
}
func changeBook2(book *Books) {
   book.title = "book1_change"
}

func main() {
    var book1 Books
    book1.title = "book1"
    book1.author = "zuozhe"
    book1.book_id = 1
    changeBook(book1)
    fmt.Println(book1)
    
    changeBook2(&book1)
    fmt.Println(book1)
}