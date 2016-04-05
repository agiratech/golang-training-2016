package main 
 
import "fmt"


type rect struct{
  length , breadth int 
}

func (r *rect)rectarea() int {
  return r.length * r.breadth 
}

func main() {

  var r rect
  r.length= 10;r.breadth=20
  fmt.Println(r.rectarea())
}