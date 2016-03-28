package main

import "fmt"

const Pi = 3.14

func main() {
	const name = "arun"
	const Pi = "Birthday"
	fmt.Println("Hello", name)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
