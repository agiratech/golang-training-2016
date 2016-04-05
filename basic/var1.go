package main

import "fmt"

func main() {
   
    var a int = 0 
    var x [5] int
    x[0] = 0
    x[1] = 1
    x[2] = 2
    x[3] = 3
    x[4] = 4 
    for i := 0;i < 5;{
    a = a + x[i]
    fmt.Println(a) 
    i = i + 1;
}
   
}