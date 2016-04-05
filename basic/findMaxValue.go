package main

import "fmt"

func main(){
	var a=10;
	var b=20;
	var res=max_Value(a,b);
	fmt.Println("The max value of 10,20 is:",res)
	
	}
	
	func max_Value(x, y int) int {
		var res int
		if (x > y) {
			res = x
			} else {
			res = y
			}
		return res
		}
	

