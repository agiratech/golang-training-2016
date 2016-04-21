package main

import (
//  "log"
  "./arunpackage"
  //"./insertsql"
  "net/http"
  "fmt"
  "strconv"
  "log"
    "github.com/gorilla/mux"
)

//var out string

  func main() {
 //out = insertsql.Db_conn()
  	//arunpackage.Db_con()
 m:=mux.NewRouter().StrictSlash(true)
  m.Handle("/", http.FileServer(http.Dir("stat")))
  //m.Handle("/", http.FileServer(http.Dir("insert")))
  //http.HandleFunc("/arunpackage",handler)
  m.HandleFunc("/arun",handler_search) 
  //m.HandleFunc("/arun/insertsql",handler_insert) 
  log.Fatal(http.ListenAndServe(":4000", m))
}

//func handler_insert(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(out))
//}

func handler_search(w http.ResponseWriter,r *http.Request) {

	value := r.FormValue("name")
    value1, _:= strconv.Atoi(value)
    search:= arunpackage.Db_con(value1)
    fmt.Fprintf(w, search) 
}