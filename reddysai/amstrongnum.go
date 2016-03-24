package main
import "fmt"
	var i,num,sum,rem,temp int
func main() {
	
	fmt.Println("Please enter the range for finding amstrong numbers")
	fmt.Scanf("%d",&num)
	for i=1;i<=num;i++ {
			temp=i
			sum = amstrong(temp)
			if(sum==i) {
				fmt.Println(i)
				}
				sum=0
			}
	}
	func amstrong(temp int) int {
				rem=temp%10
				sum+=rem*rem*rem
				temp/=10
				if(temp!=0) {
				amstrong(temp)
				}
				return sum
				}


    	


