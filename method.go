package main

import "fmt"

type Rectangle struct {
   length ,breadth int
}


func(r Rectangle) area() int {
   return r.length * r.breadth
}

func(r Rectangle) perimeter() int {
	return 2*(r.length + r.breadth)
}

func main(){
   r1 := Rectangle{4, 5}
   fmt.Println("Area : ", r1.area())
   fmt.Println("Perimeter :", r1.perimeter())

}