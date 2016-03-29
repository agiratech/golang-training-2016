package main

import "fmt"

func main() {
   var n [10]int
   var i int

    
   for i = 0; i < len(n); i++ {
      n[i] = i + 100 
       fmt.Printf("Element[%d]=%d\n",i,n[i] )
   }

}