package main

import "fmt"

type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
   var book1 Books
   var book2 Books

   // book1 描述
   book1.title = "Go 语言"
   book1.author = "www.runoob.com"
   book1.subject = "Go 语言教程"
   book1.book_id = 6495407
   
   // book2 描述
   book2.title = "Python 语言"
   book2.author = "www.runoob.com"
   book2.subject = "Python 语言教程"
   book2.book_id = 6495000

   /* 打印 Book1 信息 */
   printBook(&book1)
   
   /* 打印 Book2 信息 */
   printBook(&book2)
}

func printBook(book *Books) {
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}