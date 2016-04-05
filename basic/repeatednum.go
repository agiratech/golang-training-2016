package main
import "fmt"

func main() {
	var a[] int
	var i,j,num int
	var count int =0
	rptnum := make(map[int]int)

	
	fmt.Printf("How many elements you want to enter into the array : ")
	fmt.Scanf("%d", &num)
	fmt.Println("Please enter array values")
	a = make([]int,num)
	for i=0;i<num;i++ {
		fmt.Scanf("%d",&a[i])
			}
	
	fmt.Printf(" Repeating elements are\n ");
		for i = 0; i < num; i++ {
			for j = i+1; j < num; j++ {
				if(a[i] == a[j]) {
						count++
					rptnum[a[i]]=count+1
					} 
					}
				
					}
	for rpt:= range rptnum {
		fmt.Println("a[",rpt,"] is repeated ",rptnum[rpt])
			}
	}
