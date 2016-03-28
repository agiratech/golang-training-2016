package main 

import "fmt"

type rect struct{
	x, y int
}

func area(r *rect, i int){
	r.x = r.x + i
	r.y = r.y + i
}
func (r *rect) area(i int){
	r.x = r.x + i
	r.y = r.y + i
}

func main(){

q := rect{2,3}
area(&q,5)
fmt.Println(q)

p := rect{7,8}
p.area(5)
fmt.Println(p)

}