package main 

import "fmt"


type rect struct{
	x,y int
}


func (r *rect) area()int{
	return r.x * r.y
}

func (r *rect) scale(i int){
	r.x = r.x + i
	r.y = r.y + i
}

func main(){

a := rect{2,3}
fmt.Println(".....",a.area())
a.scale(5)
fmt.Println(".....",a.area())



}
	
