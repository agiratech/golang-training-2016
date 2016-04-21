package main

import (
//  "log"
  "./arunpackage"
  "net/http"
)

var out string
   func main() {
   	out = arunpackage.Db_con()
  //csvsql.Db_con()
  http.Handle("/", http.FileServer(http.Dir("static")))
  http.HandleFunc("/arunpackage",handler) 
  http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(out))
}
