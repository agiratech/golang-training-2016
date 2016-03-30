package main

import "fmt"

type geometry interface{
	area() float64
	peri() float64
}

type rect struct{
	l ,b float64
}

type circle struct{
	radius float64
}

func (r rect) area() float64{
	return r.l * r.b
}

func (c circle) area() float64{
	return 3.14 * c.radius *c.radius
}

func (r rect) peri() float64{
	return 2 * r.l * r.b
}

func (c circle) peri() float64{
	return 2 * 3.14 * c.radius
}

func measure(g geometry){

    area := g.area()
	fmt.Println("area : ",area)

	peri := g.peri()
	fmt.Println("perimeter : ",peri)
}
func main(){

r := rect{10,20}
c := circle{3.0}
measure(r)
measure(c)	
}