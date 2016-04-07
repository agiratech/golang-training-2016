/*#here is your sample programs to go..*/
package main

import (
	"fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func main() {
    router := httprouter.New()
    http.HandleFunc("/",handler)

    router.GET("/", Index)

    log.Fatal(http.ListenAndServe(":8088", router))
    log.Fatal(http.ListenAndServe(":3000",nil))
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
 }
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello ")
	
}
