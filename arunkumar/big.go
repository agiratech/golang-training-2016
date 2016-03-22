package main

import "fmt"

type Rectangle struct {
   length ,breadth int
}

type Books struct {
	
	id int
	title string
	author string 
	subject string
}

func(r Rectangle) area() int {
   return r.length * r.breadth
}

func(r Rectangle) perimeter() int {
	return 2*(r.length + r.breadth)
}

func printbook(book *Books) {

fmt.Println(".......",book.title)
fmt.Println(".......",book.id)
fmt.Println(".......",book.author)
fmt.Println(".......",book.subject)

}

func swap(x int, y int) (int, int){

	var temp int

    temp = x
    x = y
    y = temp 

    return x, y 

}

func main() {
	var a Books
    var b Books
    var c int = 100
	var d int = 200
    var n [2]int


    fmt.Println("Before Swapping",c)
    fmt.Println("Before Swapping",d)

    n[0],n[1]= swap(c,d)

    fmt.Println("After Swapping",n[0])
    fmt.Println("After Swapping",n[1])

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

   r1 := Rectangle{4, 5}
   fmt.Println("Area : ", r1.area())
   fmt.Println("Perimeter :", r1.perimeter())

}

