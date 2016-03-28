package main

import (
       "fmt"
       "math"
)
type student struct{
	name string
	id int
}

func squareroot(x float64) float64{
        fmt.Println(math.Sqrt(x)) 
       return x
}

func add(x , y int)int{
	return x + y
}

func swapnumbers(x , y int) (int,int){
	return y , x
}

func swapstring(x , y string) (string,string){
	return y , x
}

func main(){
    var stu[5] student
    var num ,i,x,y int
    var n float64
    var a,b string
    fmt.Println("Enter the total number of students")
    fmt.Scanf("%d",&num)
    for i = 0 ; i < num ; i++ {
	fmt.Println("Enter the name")
	fmt.Scanf("%s",&stu[i].name)
	fmt.Println("Enter the id")
	fmt.Scanf("%d",&stu[i].id)
	 fmt.Println("Enter the number")
    fmt.Scanf("%f",&n)
    fmt.Println("Enter the number for x ")
    fmt.Scanf("%d",&x)
    fmt.Println("Enter the number for y ")
    fmt.Scanf("%d",&y)
    fmt.Println("Enter the number for a ")
    fmt.Scanf("%s",&a)
    fmt.Println("Enter the number for b")
    fmt.Scanf("%s",&b)
}
for i = 0 ; i < num ; i++ {
	fmt.Println(stu[i].name)
	fmt.Println(stu[i].id)
	squareroot(n)
	fmt.Println(add(x , y))
    fmt.Println(swapnumbers(x,y))
    fmt.Println(swapstring(a,b))
	
 }
}