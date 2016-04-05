package main

import "fmt"

type student struct{

 name string
 Rollno int
 totalsubject int
 marks[10] int 

} 

func calculation() student {
var stu student
var sum, i int
     fmt.Println("Enter The Number of Subjects")
     fmt.Scanf("%d",&stu.totalsubject)
	for i = 0 ; i < stu.totalsubject ; i++ {
	fmt.Println("Enter The marks")
     fmt.Scanf("%d",&stu.marks[i])
     sum = sum + stu.marks[i]
	}
	fmt.Println("Your sum of Marks", sum)
	return stu
}

func main() {
var stu student
var totalstudents ,i int
fmt.Println("Enter the total no of students")
fmt.Scanf("%d",&totalstudents)
for i = 0 ; i < totalstudents ; i ++ {
     
    fmt.Println("Enter the Student name")
    fmt.Scanf("%s",&stu.name)
	fmt.Println("Enter the Rollno")
	fmt.Scanf("%d",&stu.Rollno)
	calculation()
}

}

