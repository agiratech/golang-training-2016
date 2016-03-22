package main

import "fmt"

var temp = 0

func main(){
	
	for j := 1; j <= 100; j++{

	if ( j%3 == 0) || ( j%6 == 0) || ( j%9 == 0) {
	  
	  temp += j
	} 

    }

    fmt.Println("Number:",temp)
}