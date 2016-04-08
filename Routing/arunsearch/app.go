package main

import (
//  "log"
  "./arunpackage"
  "net/http"
  "fmt"
  "strconv"
  
    //"github.com/gorilla/mux"
)

//var out string

  func main() {
 //out = arunpackage.Db_con()
  	//arunpackage.Db_con()
  http.Handle("/", http.FileServer(http.Dir("stat")))
  //http.HandleFunc("/arunpackage",handler)
  http.HandleFunc("/arun",handler_search) 
  http.ListenAndServe(":3001", nil)
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(out))
//}

func handler_search(w http.ResponseWriter,r *http.Request) {

	value := r.FormValue("name")
    value1, _:= strconv.Atoi(value)
    search:= arunpackage.Db_con(value1)
    fmt.Fprintf(w, search) 
}
