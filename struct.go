package main

import "fmt"

type Books struct {
	
	id int
	title string
	author string 
	subject string
}

func printbook(book *Books) {

fmt.Println(".......",book.title)
fmt.Println(".......",book.id)
fmt.Println(".......",book.author)
fmt.Println(".......",book.subject)

}

func main() {
	var a Books
    var b Books

a.id = 101
a.title = "C Programming"
a.author = "Balagurusamy"
a.subject = "Programming"

b.id = 102
b.title = "C++ Programming"
b.author = "Balagurusamy"
b.subject = "Codings"

printbook(&a)

printbook(&b)

fmt.Println(".......",&a)
fmt.Println(".......",&b)

}

