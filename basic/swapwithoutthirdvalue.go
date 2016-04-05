package main

import "fmt"


func swap(x ,y int) (int ,int){
    x = x + y
	y = x - y
	x = x - y

	return x,y
}

func main(){
	
	var a , b int
	fmt.Println("Enter the values")
	fmt.Scanf("%d",&a)
	fmt.Scanf("%d",&b)
    fmt.Println("Before Swapping : ", a , b)
	a ,b = swap(10,20)
    fmt.Println("after Swapping : ", a , b)

        
}