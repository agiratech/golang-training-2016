package main

import "fmt"

func main() {
	
	var a int = 10
	var b int = 20
	var c int = 30
    var ret int

	ret = add(a, b)
	fmt.Println("a + b : ",ret)

    ret = mul(a,b,c)
    fmt.Println("a * b * c : ",ret)
}

func add(x int,y int) int {
	
	return x+y
}

func mul(x int,y int,z int) int {
	
	return x*y*z 
}