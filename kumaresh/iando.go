package main

import "fmt"

func main() {
  fmt.Println("Enter a Details: ")
  for i := 0; i < 2; i++ {
  
  fmt.Print("Enter a Name: ")
  var name string
  fmt.Scanf("%s", &name)

  fmt.Print("Enter a Age: ")
  var age int
  fmt.Scanf("%d", &age)

  fmt.Print("Enter a Place: ")
  var place string
  fmt.Scanf("%s", &place)
  }

   // for i := 0; i < 2; i++ {
/*   fmt.Println("Student:")
   fmt.Println("Name: "+name)
   fmt.Println("Age:"+age)
   fmt.Println("Place:"+place)/**/

}