package main

import "fmt"

func main() {
   
   var  balance = []int {10, 20, 30, 40, 50}
   var avg float32


   avg = getAverage( balance, 5 ) ;


   fmt.Printf( "Average value is: %f ", avg );
}

func getAverage(arr []int, size int) float32 {
   var i int 
   var sum int = 0
   var avg float32  

   for i = 0; i < size;i++ {
      sum =sum + arr[i]
      avg = float32(sum / size)

   }
   return avg;
}
