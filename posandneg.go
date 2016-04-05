package main

import "fmt"

func main() {
 
 var sub,i,mark,neg,pos int
 fmt.Println("Enter the total number of subjects")
 fmt.Scanf("%d",&sub)
 fmt.Println("Enter the subject Marks")
for i = 0 ; i < sub ; i++ {
fmt.Scanf("%d",&mark)
if(mark < 0) {
	neg = neg + mark
} else if (mark > 0) {
	pos = pos + mark
}
}
fmt.Println("sum of all the Negative subject Marks is",neg)
fmt.Println("sum of all the Positive subject Marks is",pos)
}