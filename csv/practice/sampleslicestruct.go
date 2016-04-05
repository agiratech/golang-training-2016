package main
 

 import "fmt"
type s struct{
 name string
 num int
 }

 func main() {
  var sam []s
 var sam1 s
 sam1.name="a"
 sam1.num=1;
 sam=append(sam,sam1)
 sam1.name="b"
 sam1.num=2
 sam=append(sam,sam1)
  	

  for j:=range sam {
  	fmt.Println(sam[j].name)
  	fmt.Println(sam[j].num)
  	
  }
 }