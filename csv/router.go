package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strconv"
    "./searchdynamic"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
   // fmt.Fprint(w, "Welcome!\n")
    id :=r.FormValue("empid")
    s,_:= strconv.Atoi(id)

    js := search.Search(s)
    w.Write([]byte(js))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("reddy"))
}

func main() {
    router := httprouter.New()
    //Its serves static files
    router.NotFound = http.FileServer(http.Dir("static"))
    router.GET("/empid", Index)
    router.GET("/hello/:reddy", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}