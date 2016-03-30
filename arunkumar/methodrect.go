package main

import 
	"fmt"

type rect struct {
	X, Y int
}

func area(v rect) int {
	return v.X*v.Y
}

func main() {
	v := rect{3, 4}
	fmt.Println(area(v))
}