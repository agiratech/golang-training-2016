package main 

import "fmt"

type Stack struct {
	 
	 top *Element
	 size int
}

type Element struct {
	
	value interface{}
	next *Element
}

func (s *Stack) Len() int{
	return s.size
}

func (s *Stack) Push(value interface{}) interface{} {
	
	s.top = &Element{value, s.top}
	//fmt.Println(s.top)
	s.size++
	return s.top
}

func main() {
	
	stack := new(Stack)

	stack.Push("arun")
	stack.Push("heloo")
	
		fmt.Println(stack.value)
     
	/*for stack.Len() > 0 {

		fmt.Println(*stack)

	}
*/
}

