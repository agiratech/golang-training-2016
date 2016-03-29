package main  

import ("container/list"
        "fmt"
)

func main(){
	
	l := list.New()
	fmt.Println(l.Len())

	
	
	l.PushFront("u")
	l.PushFront("r")
	l.PushFront("a")
   
    l.PushBack("k")
    l.PushBack("u")
    l.PushBack("a")
    l.PushBack("r")

	

	for e := l.Back() ; e!= nil ; e = e.Prev(){

fmt.Print(e.Value)

    }	
    for e := l.Front() ; e!= nil ; e = e.Next(){ 

	 fmt.Println(e.Value)
	}
	fmt.Println(l.Len())
}