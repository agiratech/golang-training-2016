package main

import "fmt"

func main() {
  
  var n int
  temp := 0
  temp2 := 0
  fmt.Scanf("%d", &n)
  
  for i := 1; i <= n; i++ {

    if i % 2 == 0 {
       fmt.Println("even:",i)
       temp += i
       
    } else {
      
       fmt.Println(" odd:",i)
       temp2 += i
    
    }

  }
  fmt.Println("Even NO/- total:%v",temp)
  fmt.Println("Odd NO/- total:%v",temp2)
}