package main

import "fmt"

type rect struct {
	x,y int
}

func area(r rect) int{
	return r.x * r.y
}

func areareplica(r *rect){
	r.x = 2
	r.y = 4 
}

func areareplica1(r *rect){
	r.x = 50
	r.y = 40
}

func (r *rect) areareplica2(){
	r.x = 7
	r.y = 4
}

func areareplica3(r *rect , i int){
	r.x = r.x + i
	r.y = r.y + i
}

func (r *rect) areareplica4(i int){
	r.x = r.x + i
	r.y = r.y + i
}


func main(){

r := rect{10,20}
	fmt.Println(area(r))
	areareplica(&r)
	fmt.Println(area(r))
	areareplica1(&r)
    fmt.Println(area(r))
    r.areareplica2()
    fmt.Println(area(r))
    areareplica3(&r,7)
     fmt.Println(area(r))
      r.areareplica4(20)
      fmt.Println(area(r))
}