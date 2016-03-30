package main

import (
      "fmt"
      "math"
)	

func main(){
	
	x,y := 3,4
	z:=math.Sqrt(float64(x*x+y*y))
	fmt.Println(x,y,z)
}