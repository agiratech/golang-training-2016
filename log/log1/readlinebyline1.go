package main

import (
  "bufio"
  "log"
  "os"
  "fmt"
)

func checkerror(err error){

  logfile,err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
 //if err!=nil{
  // panic(err)
   // }
  //defer logfile.Close()
  log.SetOutput(logfile)
  log.Println("some error in the current code")

 } 

func logg(){

  f,_:=os.Open("log.txt")
  scanner := bufio.NewReader(f)
   s,_:= scanner.ReadString(10)
  fmt.Println(s)
}

func main() {

  // open a file
    file, err := os.Open("hell.txt")
   // checkerror(err)
     if err!=nil {
         checkerror(err)
         logg()
        }
      
    // make sure it gets closed
    //defer file.Close()

    // create a new scanner and read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        log.Println(scanner.Text())
    }

    // check for errors
   // err = scanner.Err()  
    //if err != nil {
      //log.Fatal(err)
    //}

}