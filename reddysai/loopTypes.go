package main
import "fmt"

func main() {
	var num [4]int
	var i, j int 
	for i=0;i<len(num);i++ {
		fmt.Scanf("%d", &num[i])
		}
		i = 0
	/*for i<len(num) {
		fmt.Println(num[i])
		i++
		}*/
	fmt.Println("Before Swapping")
	for k,x := range num {
		fmt.Printf("num[%d] = %d\n", k, x)
		}

	for i=0;i<len(num);i++ {
		for j=0;j<=i;j++ {
			if(num[i] < num[j]) {
				t:=num[i]
				num[i]=num[j]
				num[j]=t
					}
				}
			}
	fmt.Println("After Swapping")
	i = 0
	for i<len(num) {
		fmt.Printf("num[%d] = %d\n", i, num[i])
		i++
		}
	}
