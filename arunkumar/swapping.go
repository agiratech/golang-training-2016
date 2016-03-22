package main

import "fmt"
    

func main(){

    var a int = 100
	var b int = 200
    var ret int 
    var re int

    fmt.Println("Before Swapping",a)
    fmt.Println("Before Swapping",b)

    ret,re = swap(a,b)

    fmt.Println("After Swapping",ret)
    fmt.Println("After Swapping",re)
}

func swap(x int, y int) (int, int){

	var temp int

    temp = x
    x = y
    y = temp 

    return x, y 

}