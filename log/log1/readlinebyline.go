package main

import (
  "bufio"
  "log"
  "os"
)


func checklog(err error){

  if err!= nil {
      logfile,_:= os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
      log.SetOutput(logfile)
      log.Println("some error in your file")
      //log.Panic(err)
     // log.Fatal(err)
     f,_:= os.Open("log.txt")
      scanner:= bufio.NewScanner(f)
      //for scanner.Scan() {
        log.Println(scanner.Text())
    //}

    }
}


func main() {

  // open a file
    file,err:= os.Open("hell.txt")
    checklog(err)
    //if err!= nil {
      //logfile,_:= os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
     // log.SetOutput(logfile)
     // log.Println("some error in your file")
   // log.panic(err)
    //}
    // make sure it gets closed
  //  defer file.Close()
    
    // create a new scanner and read the file line by line
    scanner:= bufio.NewScanner(file)
   for scanner.Scan() {
        log.Println(scanner.Text())
    }
   //checklog(err)
    // check for errors
    // err := scanner.Err()
     // if err!= nil {
     // log.Fatal(err)
    // } 

}