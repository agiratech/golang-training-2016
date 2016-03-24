package main

import "fmt"

func main() {
 
 var sub,sum,i int
 var marks[10] int
 fmt.Println("Enter the total number of subjects")
 fmt.Scanf("%d",&sub)
 fmt.Println("Enter the subject Marks")
for i = 0 ; i < sub ; i++ {
fmt.Scanf("%d",&marks[i])
sum = sum + marks[i] 
}
fmt.Println("sum of all the subjects is",sum)	
}