package main
import ("fmt"; "time")
var n int
func main() {
	
	n=10
	go process(n)
	go func() {
		time.Sleep(time.Millisecond *10)
				fmt.Println("",n*n*n/2*n)

	}()
	go func() {
		fmt.Println("",n*n)
		time.Sleep(time.Millisecond * 10)
		fmt.Println("",n/2)
	}()
	func() {
	fmt.Println(n)
	time.Sleep(time.Millisecond * 100)
	}()
}
func process(n int){
	fmt.Println((n*n)/2)
	time.Sleep(time.Second * 10)
}