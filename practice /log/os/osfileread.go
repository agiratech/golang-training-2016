package main 

import ("fmt"
         "os"
         "log")

func main() {
	
	file,err:= os.Open("hello.txt")
	if err != nil{
	 //fmt.Println("some error")
	log.Fatal(err)
	}
    data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %s \n", count, data)
}