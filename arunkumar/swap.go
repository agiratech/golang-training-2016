package main

 import "fmt"

 func main(){

 var a int = 100;
 var b int = 200;

 fmt.Println("Before Swapping",a)
 fmt.Println("Before Swapping",b)
 
 a,b = swap(a,b)

 fmt.Println("After Swapping",a)
 fmt.Println("After Swapping",b)
 
 }

func swap(x , y int) (int, int) {

	return y, x ;


}