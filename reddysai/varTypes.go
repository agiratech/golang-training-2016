package main
import "fmt"

func main() {
	var x int
	var y float32
	var z string
	var a, b, c = 10, 22.22, "reddy"
	x = 10/2
	y = 32.12-0.12
	z = "sai"

	i := "welcome to AgiraTech"
	j := 30.00
	k := 15

	fmt.Printf("a value is : %d\t\t and Datatype is : %T\n", a,a)
	fmt.Printf("b value is : %f\t\t and Datatype is : %T\n", b,b)
	fmt.Printf("c value is : %s\t\t and Datatype is : %T\n", c,c)
	fmt.Printf("x value is : %d\t\t and Datatype is : %T\n", x,x)
	fmt.Printf("y value is : %f\t\t and Datatype is : %T\n", y,y)
	fmt.Printf("z value is : %s\t\t and Datatype is : %T\n", z,z)
	fmt.Printf("i value is : %s\t\t and Datatype is : %T\n", i,i)
	fmt.Printf("j value is : %f\t\t and Datatype is : %T\n", j,j)
	fmt.Printf("k value is : %d\t\t and Datatype is : %T\n", k,k)
	}
