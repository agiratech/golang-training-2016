package main

import "fmt"

func swap(x ,y string) (string, string) {
	return y, x	
}

func main() {
	a, b := swap("Arun", "kumar")
	fmt.Println(a,b)
}