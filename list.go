package main

import ( "container/list"
         "fmt"
)

func main() {
	l := list.New()
	e4 := l.PushBack(2)
	e1 := l.PushFront(4)
	l.InsertBefore(1, e4)
	l.InsertAfter(3, e1)

	for e:= l.Front(); e != nil; e = e.Next() {
	fmt.Println(e.Value)
	}
}