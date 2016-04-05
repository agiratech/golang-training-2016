package main

import "fmt"

func main() {
	var i,greatest int
	var num =[]int{2,8,5,1,2}
	greatest = num[0]
	for i = 1;i < len(num);i++ {
		if(num[i] > greatest ) {
			greatest = num[i]
		}
	}	
	fmt.Println("Greatest Number : ",greatest)  
} 