package main

import "fmt"

func zero(x *int){
	*x = 0
}

func main() {
	x := 5 
	z := 51
	zero(&x)
	fmt.Println(x)
	fmt.Println("The address of the variable", &z)
    fmt.Println("The value in the variable", z)

}