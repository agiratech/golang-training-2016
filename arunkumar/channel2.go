package main

import ("fmt"
       "time")

var counter int = 0
func main(){
	
	
    for i :=0; i<5 ; i++{
    time.Sleep(time.Second*3)
    go process()
    } 
    time.Sleep(time.Second*10)
    }
    func process(){
    counter++
    fmt.Println(counter)
    }
