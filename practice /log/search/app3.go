package main

import (
//  "log"
  "./arunpackage"
  "net/http"
  "log"
   "strconv"
   //"time"
   "os"
   "fmt"
   "bufio"
  //"io/ioutil"

)

var logFile *os.File 

type Options struct {
  Path string
  Port string
}

//func checking(err error) {

  //logfile,err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND,0666)
  //log.SetOutput(logfile)
  //log.Println("some error in the current code")

 //} 

 //func logg() {

 // f,_:=os.Open("log.txt")
 // scanner := bufio.NewReader(f)
  // s,_:= scanner.ReadString(10)
  //fmt.Println(s)
//}
   
func main() {


 // var err error
   // logFile, err = os.Create("logfile.txt")
  //  if err != nil {
   //     log.Fatal("Log file create:", err)
   //     return
  //  }
  //  defer logFile.Close()
  http.HandleFunc("/arunpackage",handler_search) 
  http.Handle("/", http.FileServer(http.Dir("static")))
  err = http.ListenAndServe(":3002", Log(http.DefaultServeMux))
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      s,err := os.OpenFile("log.txt")
   s := fmt.Printf("%s %s %s %s\n", time.Now r.RemoteAddr, r.Method, r.URL.Path[1:], r.Proto, r.Host)
  log.SetOutput(s)
    handler.ServeHTTP(w, r)
  })
}

func handler_search(w http.ResponseWriter,r *http.Request) {

	value := r.FormValue("name")
    value1, _:= strconv.Atoi(value)
    search:= arunpackage.Db_con(value1)
    fmt.Fprintf(w, search) 
}
