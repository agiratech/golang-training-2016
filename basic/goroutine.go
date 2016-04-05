package main

import (
    "fmt"
    "time"
)
var counter int =0
func main() {
    for i:=0; i<5;i++{
    go process()
    process()
}
    time.Sleep(time.Millisecond * 10)
   
}
func process(){
    counter++
    fmt.Println(counter)
}