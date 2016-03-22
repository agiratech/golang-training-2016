package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20
	var ret int
    ret = max(a,b)

    fmt.Println("The Maximum value is",ret)
}

func max(num1, num2 int) int {

        var result int

	    if(num1<num2){
        result = num1
		}else{
	    result = num2
		}
		return result;
}