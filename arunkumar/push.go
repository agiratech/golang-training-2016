package main

import "fmt"

type Stack struct {

	top *Element

	size int
}

type Element struct{

	value interface{}

	next *Element
}


func (s *Stack) Push(value interface{}) {

	s.top = &Element{value, s.top}

	s.size++

	return
}


func (s *Stack) Pop(value interface{}) {
		
			if s.size > 0 {
			
			value, s.top = s.top.value, s.top.next
			
			s.size--
			
			return
	}

	return
}


func main() {

	var n int

	stack:= new(Stack)

    var st[3]string
 
	fmt.Println("Enter your choice")

	fmt.Println("Enter 1 for Push")

	fmt.Println("Enter 2 for pop")  

	fmt.Scanf("%d",&n)

     switch(n){

     case 1 :

	    fmt.Println("Push: ")

		for i :=0; i < 3; i++ {

				fmt.Scanf("%s", &st[i])

	    		stack.Push(st)
        }

	case 2 :

		fmt.Println("Pop: ")

		for i := 0 ; i < 3 ; i++ {

         stack.Pop(st)

		}

    default :

		fmt.Println("Enter the correct value")

		}

}