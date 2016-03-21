package main
import "fmt"

func main() {
		var values = []int {10, 22, 07, 44}
		var res = float32(avg_Value(values, 4))
		fmt.Printf("The average value is: %f\n", res)
	}
func avg_Value(values[]int, length int) float32 {
			var i,sum int
			for i=0; i<length; i++ {
				sum += values[i]
				}
			var res = float32(sum/length)
			return res;
	}
