package main

import (
	"fmt"
	
)

type rect struct {
	x, y int
}

func area(r *rect) int {
	return r.x * r.y
}

func (r *rect) Scale() int {
	r.x = 10
	r.y = 20
	return r.x * r.y
}

func main() {
	r := rect{3, 4}
	fmt.Println(&area(r))
	fmt.Println(r.Scale())
	fmt.Println(&area(r))
}
