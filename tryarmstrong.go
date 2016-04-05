package main

import "fmt"

 func armstrong(n int) int {
   
   var temp int = n
   var remainder,sum int
   //var i int = 0
   {
      remainder = temp%10
      temp = temp/10
      sum = sum + (remainder*remainder*remainder)
   }
  return sum
      }

func main () {
   var n,number int
 
   fmt.Println("Enter Max range: ")
       fmt.Scanf ("%d", &n)
      fmt.Println("Printing armstrong numbers from 0 - %d\n", n)
      for number = 0; number <= n; number++ {
         if(armstrong(number)==n) {
          fmt.Println(number)
         }
      }
     } 
