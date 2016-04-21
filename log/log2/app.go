package main

import (
//  "log"
  "./arunpackage"
  "net/http"
  "fmt"
  "log"
   "strconv"
   "time"
   "encoding/json"
  //"io/ioutil"

)


type Options struct {
  Path string
  Port string
}

//var out string
   
func main() {
  	//out = arunpackage.Db_con()
  //csvsql.Db_con()
  //http.Handle("/", http.FileServer(http.Dir("static")))
  http.HandleFunc("/arunpackage",handler_search) 
  //http.ListenAndServe(":3000", nil)


//}
 // op := &Options{Path: "./", Port: "8080"}

  //data, _ := ioutil.ReadFile("./config.json")

  //json.Unmarshal(data, op)

  http.Handle("/", http.FileServer(http.Dir("static")))
  err := http.ListenAndServe(":3002", Log(http.DefaultServeMux))
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}

func Log(handler http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%s %s %s %s %s\n",time.Now(),r.RemoteAddr, r.Method, r.URL.Path[1:], r.Proto, r.Host)
    handler.ServeHTTP(w, r)
  })
}

func handler_search(w http.ResponseWriter,r *http.Request) {
  byte, _ := json.Marshal(r)
  fmt.Println()
	value := r.FormValue("name")
    value1, _:= strconv.Atoi(value)
    search:= arunpackage.Db_con(value1)
    fmt.Fprintf(w, search) 
}
