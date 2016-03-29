package main

import "fmt"

var i int
var n int

type student struct {
var  Name string
var  RollNo int
var  Subjects int
var  OverallMarks int 
var average int
}

func studentdetails() Student {
fmt.Println("Enter Number of Students")
fmt.Scanf("%d", &n)
for i = 0 ; i < n ; i++ {
fmt.Println("Enter a Student number Details")
fmt.Println("Enter Student Name :")
fmt.Scanf("%s", &Name)
fmt.Println("Enter Student Roll-No : ")
fmt.Scanf("%d", &RollNo)
fmt.Println("Enter Student Subjects : ")
fmt.Scanf("%d", &Subjects)
fmt.Println("Enter Student Over-all_Marks : ")
fmt.Scanf("%d", &OverallMarks)

average =  OverallMarks / Subjects

fmt.Println("Your average : ",average)

}
return Student
}

func main() {
 
 var a Student
 a =studentdetails()
}