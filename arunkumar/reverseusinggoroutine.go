package main

import ("fmt"
       "time"
 )

var rem int
var rev = 0

func reverse(n int)int{
	
	for n > 0 {

	rem = n%10
    rev = rev * 10 + rem
	n = n/10
} 
time.Sleep(100 * time.Millisecond)
return  rev
}

 func reverse1(value string) string {
  
    data := []rune(value)
    result := []rune{}

    
    for i := len(data) - 1; i >= 0; i-- {
	result = append(result, data[i])
    }
time.Sleep(100 * time.Millisecond)
    return string(result)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){

	var num int
	var value1 string
	fmt.Println("Enter the numbers")
	fmt.Scanf("%d",&num)
    rem = reverse(num)
    fmt.Println(rem)
time.Sleep(1000 * time.Millisecond)
    fmt.Println("Enter the string")
	fmt.Scanf("%s",&value1)
    reversed1 := go reverse1(value1)
    fmt.Println(reversed1)

   time.Sleep(1000* time.Millisecond)

    go say("world")
	say("hello")


}