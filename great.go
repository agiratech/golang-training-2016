package main

import "fmt"

func main() {
	        var inNumber int = 100
            var largest int
	        var i int = 0
	        while(i >= 0){	
	         fmt.Println("Please enter a bunch of values. To end, enter a negative number.")
                if(inNumber > largest) {
	                largest = inNumber
	            }
	        }   
	        return largest
	        fmt.Println("The largest input value was ",largest)
	    }  
