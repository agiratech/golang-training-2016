package main

import "fmt"

func main() {
 
 var i,tot,sum int
 var number[10] int
 fmt.Println("Enter the Total number")
 fmt.Scanf("%d",&tot)
 sum = number[0]
 for i = 0 ; i < tot ; i++ {
 	fmt.Scanf("%d",&number[i])
if(sum < number[i]){
	sum = number[i] 
}
}
fmt.Println("Your Greatest Number is",sum)
 }