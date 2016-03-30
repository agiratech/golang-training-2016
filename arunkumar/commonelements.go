package main

import "fmt"

func main() {
	var i,j,n,z int
	var x []int
	var y []int
	fmt.Println("Enter the x value")
	fmt.Scanf("%d",&n)
    x = make([]int,n)
    fmt.Println("Enter the x Elements")
	for i = 0 ; i < n ; i++{
		fmt.Scanf("%d",&x[i])
	}
	fmt.Println("Enter the y value")
	fmt.Scanf("%d",&z)
	y = make([]int,z)
	fmt.Println("Enter the y Elements")
	for j = 0 ; j < z ; j++{
		fmt.Scanf("%d",&y[j])
	}
	for i = 0 ; i<len(x); i++ {
    	for j = 0 ; j<len(y); j++ {
    					if(x[i]==y[j]) {
    							fmt.Println("The Commom Values are ",x[i])
    									}
    								}
								}
							}
						