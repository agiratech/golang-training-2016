package main
import (
	"log"
	"net/http"
	"fmt"

)

func main() {
	http.Handle("/",http.FileServer(http.Dir("practice")))
	http.HandleFunc("/empid",handler)
	log.Fatal(http.ListenAndServe(":4040",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "hello, ")
}