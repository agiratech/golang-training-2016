package main

import (
	"fmt"
	
)

type rect struct {
	x, y int
}

func area(r rect) int {
	return r.x * r.y
}

func (r *rect) Scale(f int) {
	r.x = r.x * f
	r.y = r.y * f
}

func main() {
	r := rect{3, 4}
	fmt.Println(area(r))
	r.Scale(10)
	fmt.Println(area(r))
}
