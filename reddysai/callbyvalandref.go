package main
import "fmt"

func main() {
		var x int = 10
		fmt.Println("Before calling with value the x value is :",x)
		callByValue(x)
		fmt.Println("After calling  the x value is :",x)
		fmt.Println("Before calling with reference the x value is :",x)
		callByReference(&x)
		fmt.Println("After calling the x value is :",x)
}

func callByValue(x int) {
		x+=10
		fmt.Println("in called fuction the x value is :",x)
}
func callByReference(x *int) {
		*x+=10
		fmt.Println("in called fuction the x value is :",*x)
}
