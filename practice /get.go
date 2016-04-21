package main 

import ( "net/http"
        "io")

func Hello(w http.ResponseWriter , r *http.Request){
       io.WriteString(w, "heloo world!")
}

func main() {

    mux := http.NewServeMux()
    mux.HandleFunc("/",Hello)
    http.ListenAndServe(":7800",mux)
}