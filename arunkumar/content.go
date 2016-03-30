package main

import "fmt"

type student struct{
	name string
	rollno int
	marks int
}

func main(){
	var num,i int
	p := &student{}
	a := []student{}
	fmt.Println("Enter the content")
	fmt.Scanf("%d",&num)
	
    for i = 0 ; i < num ; i++ {
    fmt.Println("Enter the Name")
	fmt.Scanf("%s",&p.name)
	fmt.Println("Enter the Rollno")
	fmt.Scanf("%d",&p.rollno)
    fmt.Println("Enter the marks")
	fmt.Scanf("%d",&p.marks)
	a = append(a,*p)
	fmt.Println(".......",a)
	}	
}