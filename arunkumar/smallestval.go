package main

import "fmt"

func main() {
	x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,2,
  19,97,17,9,
}
z := x[0]
for i := 0 ; i < len(x) ; i++ {
	if(x[i]<z) {
	z = x[i]
	}
}
fmt.Println("The smallest value is", z)
}