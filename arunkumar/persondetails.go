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
for i = 0; i < n ; i++{
fmt.Println("Enter Student Name : ")
fmt.Scanf("%s", &stu.Name)
fmt.Println("Enter Student Roll-No : ")
fmt.Scanf("%d", &stu.RollNo)
fmt.Println("Enter Student Subjects : ")
fmt.Scanf("%d", &stu.Subjects)
fmt.Println("Enter Student Over-all_Marks : ")
fmt.Scanf("%s", &stu.OverallMarks)

fmt.Println("Your Total Subjects average Marks is : ",average)
average = stu.OverallMarks/stu.Subjects
}
return stu  
}
 
func main() {
 
 var b Student 
 b = studentdetails()
  
}