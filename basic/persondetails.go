package main

import "fmt"

var n int
var i int
var average int
type Student struct{
  
  Name string
  RollNo int
  Subjects int 
  OverallMarks int
}

func studentdetails() Student {
var stu Student
fmt.Println("Enter Number of Students")
fmt.Scanf("%d", &n)
fmt.Println("Enter Student Name : ")
fmt.Scanf("%s", &stu.Name)
fmt.Println("Enter Student Roll-No : ")
fmt.Scanf("%d", &stu.RollNo)
fmt.Println("Enter Student Subjects : ")
fmt.Scanf("%d", &stu.Subjects)
fmt.Println("Enter Student Over-all_Marks : ")
fmt.Scanf("%d", &stu.OverallMarks)
average = stu.OverallMarks/stu.Subjects
fmt.Println("Your Total Subjects average Marks is : ",average)
return stu  
}
 
func main() {
 var a Student 
  a = studentdetails()
	fmt.Println("Hi ",a.Name,"(",a.RollNo,") Your average is",average,"in ",a.Subjects,"Subjects")
  
}



