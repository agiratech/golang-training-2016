package main
import "fmt"

func main() {
	var a []int 
	var n,i,j,big,position int
	fmt.Println("How many numbers You want to insert into the array")
	fmt.Scanf("%d", &n)
	a = make([]int,n)
	fmt.Println("Enter array values")
	for i = 0; i < n;i++ {
		fmt.Scanf("%d", &a[i])
			}
	/*for i:= range a {
		fmt.Println("a[",i,"] = ",a[i])
			}*/
	for i=0;i<n;i++ {
		for j=0;j<i;j++ {
			if(a[i]>a[j]) {
				big=a[i]
				position=i
					}
				}
			}
	fmt.Println("The biggest number in the given array values is :",big,"and that position :",position+1)				
		
	}
