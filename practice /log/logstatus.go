package main

import (
  "bufio"
  "log"
  "os"
  "fmt"
  "time"
  //"log/syslog"
)

func checkerror(err error) {

  logfile,err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
  log.SetOutput(logfile)
  log.Println("some error in the current code")

 } 
func logg() {

  f,_:=os.Open("log.txt")
  scanner := bufio.NewReader(f)
   s,_:= scanner.ReadString(10)
  fmt.Println(s)
}

func main() {

  // logwriter, e := syslog.New(syslog.LOG_NOTICE, "myprog")
    //if e == nil {
      //  log.SetOutput(logwriter)
    //}
    file, err := os.Open("hello.txt")
     if err!=nil {
      start := time.Now()
         checkerror(err)
         logg()
      stop := time.Now()
     fmt.Println(start,stop)
        }
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        log.Println(scanner.Text())
    }

}