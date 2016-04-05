package main

import "fmt"

type student struct{
	name string
	rollno int
	marks int
}

 func main() {

 	a :=[]student{}
    var num int
 	fmt.Println("Enter the number of students")
 	fmt.Scanf("%d",&num)
 	a = stu(num)
 	fmt.Println("......",a)

 }

 func stu(x int) []student{
    p :=&student{}
    a :=[]student{}
	for i := 0 ; i < x ; i++ {
    fmt.Println("Enter the Name")
	fmt.Scanf("%s",&p.name)
	fmt.Println("Enter the Rollno")
	fmt.Scanf("%d",&p.rollno)
    fmt.Println("Enter the marks")
	fmt.Scanf("%d",&p.marks)
	a := append(a,*p)
	fmt.Println("......",a)
 }
 return a
}