package main  

import ("container/list"
        "fmt"
)


type Element struct {

	next ,Prev *Element

	list *List


}


func (e *Element) Next() *Element {
 
 if p := e.next; e.list !=nil && p != &e.list.root {

 	return p
    }
  return nil
}

type List struct {
	
	root Element
	len int
}


func(l *List)Init() *List {   //Init initializes or clears list l
	
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *List {

	return new(List).Init()   //New returns an intialized list
}

func (l *list) LazyInit(){     // Initializes a Zero List value
  
  if l.root.next == nil {
  	l.Init()
  }

}

func (l *List) Len() int {     //Len returns the number of elements of list l

 return l.len 

 }

func (l *List) Front() *Element {      // Front returns the first element of list l or nil
	
	if l.len == 0 {
	return nil
    }
    return l.root.Next
}

func (l *List) Back() *Element {       // Back returns the last element of list l or nil
	
	if l.len == 0 {
	return nil
    }
    return l.root.Prev
}

