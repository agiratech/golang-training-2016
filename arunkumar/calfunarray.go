package main

import "fmt"
    

func main(){

    var a int = 100
	var b int = 200
    var n [2]int

    fmt.Println("Before Swapping",a)
    fmt.Println("Before Swapping",b)

    n[0],n[1]= swap(a,b)

    fmt.Println("After Swapping",n[0])
    fmt.Println("After Swapping",n[1])
}

func swap(x int, y int) (int, int){

	var temp int

    temp = x
    x = y
    y = temp 

    return x, y 

}