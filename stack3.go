package main

import "fmt"

type stack struct {
	
	stk int
	top int
}

func (s *stack) push(){
	
var num int
if(s.top == (MAXSIZE - 1)){
	fmt.Println("Stack is full")
}else {
	
	fmt.Println("Enter the element to be Pushed")
	fmt.Scanf("%d",&num)
	s.top = s.top + 1
	s.stk[s.top] = num 
}
return
}

func print() {
	
	if(s.top == -1){
	fmt.Println("stack is empty")
	}else {
	fmt.Println("The status of the stack is ")
	for i :=s.top;i>=0;i-- {
	fmt.Pritnln("%d",s.stk[i])
	}
	}
	}

	func main(){

	push()
	print()
	}
