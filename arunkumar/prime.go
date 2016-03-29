package main

import "fmt"

func main() {

   var i, j, p int

   for i=2; i < 100; i++ {
      for j=2; j <i; j++ {
         if(i%j==0) {
            p=0;
            break;
         }
         p=1
      }
      if(p==1) {
         fmt.Printf("%d is prime\n", i);
      }
   }  
}