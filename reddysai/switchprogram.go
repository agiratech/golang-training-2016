package main
import "fmt"

func main() {
	var x, y int = 5, 6
	var t int
	var c string = "y"
	
	for c =="y" || c == "Y" {
		fmt.Println("these are available operations")
		fmt.Println("1. sum\n2. mul\n3. sub")
		fmt.Println("which operation you want to do :")
		fmt.Scanf("%d",&t)
		switch t {
			case 1: fmt.Println("the sum of 5,6 is :", x+y)
			case 2: fmt.Println("the multiplication of 5,6 is :", x*y)
			case 3: fmt.Println("the substraction of 5,6 is :", x-y)
			default : fmt.Println("unknown operation please try again")
		}
		fmt.Println("if you want to continue enter 'y' :")
		fmt.Scanf("%s",&c)
	}
}