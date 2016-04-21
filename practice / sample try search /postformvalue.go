package main

import(
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
   // r.ParseForm()
	name := r.FormValue("name")
     fmt.Fprintf(w, "Helo, %s!", name)
}

func main() {

  http.Handle("/", http.FileServer(http.Dir("stats")))
  http.HandleFunc("/name",hello) 
  http.ListenAndServe(":3001", nil)

}