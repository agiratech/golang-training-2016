package main

import "fmt"

var n int
var i int
<<<<<<< HEAD
var average int
type Student struct{
=======

type Student struct {
>>>>>>> 73576cc831ecb76bc573ecc24d32ec359b155122
  
  Name string
  RollNo int
  Subjects int 
  OverallMarks int
}

func studentdetails() Student {
var stu Student
fmt.Println("Enter Number of Students")
fmt.Scanf("%d", &n)
for i = 0; i < n ; i++{
fmt.Println("Enter Student Name : ")
fmt.Scanf("%s", &stu.Name)
fmt.Println("Enter Student Roll-No : ")
fmt.Scanf("%d", &stu.RollNo)
fmt.Println("Enter Student Subjects : ")
fmt.Scanf("%d", &stu.Subjects)
fmt.Println("Enter Student Over-all_Marks : ")
fmt.Scanf("%d", &stu.OverallMarks)

fmt.Println("Your Total Subjects average Marks is : ",average)
average = stu.OverallMarks/stu.Subjects
}
return stu  
}
 
func main() {
 var a Student 
  a = studentdetails()
	fmt.Println("Hi ",a.Name,"(",a.RollNo,") You got ",a.OverallMarks,"in ",a.Subjects)
  
}
