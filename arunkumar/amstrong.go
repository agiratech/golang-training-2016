package main

import "fmt"
var sum,remainder,i int
 func armstrong(n int) int {
   
      remainder = n%10
      sum = sum + (remainder*remainder*remainder)
      n = n/10
      if(n!=0){
      	armstrong(n)
      }
         
   return sum
 }

func main () {
   var n int
 
   fmt.Println("Enter a number")
   fmt.Scanf("%d", &n)
for i = 0 ; i < n ; i++{
   sum = armstrong(i)
if (i == sum) {
      fmt.Println("Given number is an armstrong number.", i) 
   }
   sum=0
 }
}