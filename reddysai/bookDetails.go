package main
import "fmt"

type Books struct {
	title string
	author string
	cost float32
	}

func main() {
	var book_c Books
	var book_corejava Books
	var num int

	book_c.title = "C Language"
	book_c.author = "Balaguruswamy"
	book_c.cost =	390

	book_corejava.title = "Core Java"
	book_corejava.author = "R. Nageswara Rao"
	book_corejava.cost = 499
	fmt.Printf("We Provide two books...\n that are 1. C language\n\t2. Core Java \n which book You Want please enter 1 or 2:")
	fmt.Scanf("%d", &num)
	if(num == 1) {
		printBook(book_c)
		} else if(num == 2) {	
		printBook(book_corejava)
		} else {
		fmt.Println("Sorry, you Entered wrond information")
		}
	}
func printBook(book Books) {
		fmt.Println("Book Name:", book.title)
		fmt.Println("Book Author:", book.author)
		fmt.Println("Book Price:", book.cost)
	}
