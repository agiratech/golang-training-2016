package main

import "fmt"

func main() {
   
   var a int = 100
   var b int= 200

   fmt.Println("Before swap, value of a :", a )
   fmt.Println("Before swap, value of b :", b )

   
  a,b = swap(a, b)

   fmt.Println("After swap, value of a :", a )
   fmt.Println("After swap, value of b :", b )
}

func swap(x int, y int) (int, int) {
   var temp int
   temp = x    
   x = y   
   y = temp
   return x, y;     
}