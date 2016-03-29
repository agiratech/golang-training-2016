package main

import "fmt"

func main() {

  var name string
  var age int
  var place string

  fmt.Println("Enter a Details: ")
  
  
  fmt.Print("Enter a Name: ")
  fmt.Scanf("%s", &name)

  fmt.Print("Enter a Age: ")
  fmt.Scanf("%d", &age)

  fmt.Print("Enter a Place: ")
  fmt.Scanf("%s", &place)
  

   
   
   fmt.Printf("Hi %s Your age is %d and You are from %s \n", name, age, place)

}
