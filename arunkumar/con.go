package main

import ("fmt"
       "time"
)
func main(){
	  go fun(0)
	  time.Sleep(1000)
	  jump(1)
	  
}
func fun(n int) {
	time.Sleep(100)
	fmt.Println(n)
}
func jump(i int) {
	fmt.Println(i)
}