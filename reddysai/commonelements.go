package main
import "fmt"

func main() {
	var first []int
	var second []int
	var fnum,snum,i int
	fmt.Println("How many elements you want to enter in the first array")
	fmt.Scanf("%d",&fnum)
	
	first = make([]int,fnum)
	
	fmt.Println("Enter first array elements ")
	for i=0;i<fnum;i++ {
			fmt.Scanf("%d",&first[i])
			}
	
	fmt.Println("How many elements you want to enter in the second array")
	fmt.Scanf("%d",&snum)
	second = make([]int,snum)
	fmt.Println("Enter second array elements ")
	for i=0;i<snum;i++ {
			fmt.Scanf("%d",&second[i])
			}
	commonElements(first,second)
	}
func commonElements(f,s []int) {
		var i,j int
		var com,uncom []int
		for i=0;i<len(f);i++ {
			for j=0;j<len(s);j++ {
				if(f[i]==s[j]) {
					com=append(com, f[i])
						
				
						}else { 

							//uncom=append(uncom, f[i])}
					     }
					}
	fmt.Println("The common values are ")
	for i=0;i<len(com);i++ {
			fmt.Printf("%d ",com[i])
				}
	fmt.Println("\nThe uncommon values are ")
	for i:= range uncom {
			fmt.Printf("%d ",uncom[i])
			}						
	}

		
	
