package main

import "fmt"

func main() {
	var i,greatest,n int
	var num [] int
	fmt.Println("Enter the value")
	fmt.Scanf("%d",&n)
	num = make([]int,n)
	fmt.Println("Enter the elements")
	for i=0;i<n;i++{
		fmt.Scanf("%d",&num[i])
	}
	greatest = num[0]
	for i = 1;i < len(num);i++ {
		if(num[i] > greatest ) {
			greatest = num[i]
		}
	}	
	fmt.Println("Greatest Number : ",greatest)  
} 