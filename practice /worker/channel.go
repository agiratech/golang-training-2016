package main
 
import (
  "time"
  "fmt"
  "strconv"
  )

var i int

func makecakeandsend(cs chan string) {
 for i:=1; i <=3; i++ {
 cakename := "Strawberry " + strconv.Itoa(i)
 fmt.Println("Making a cake and sending")
 cs <- cakename
  }
}

func receivecakeandpack(cs chan string) {

 for i:=1; i<=3; i++ {
 s:= <-cs
 fmt.Println("Packing recevied cake", s)
  }
}


func main() {

cs:= make(chan string)

for i:=0; i < 3; i++{

	go makecakeandsend(cs)
	go receivecakeandpack(cs)
    time.Sleep(1 * time.Second)
}


}