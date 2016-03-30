package main

import "fmt"

func main() {

   arr := [5]int{1,2,3,4,5}
   x := arr[2:4]
   a := [6]string{"a","b","c","d","e","f"}
   b := a[2:5]
   slice1 := []int{1,2,3,4,5}
   slice2 := append(slice1,4,5)
   fmt.Println(x)
   fmt.Println(b)
    fmt.Println(slice1,slice2)  
  }

