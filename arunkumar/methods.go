package main

import "fmt"

type rect struct{
	x ,y int
}

func (r *rect)area() int{
	
	return r.x * r.y 
}

func main(){
 
 r := rect{2,4}
 fmt.Println(r.area())
	fmt.Println(r.x)
	fmt.Println(r.y)
}