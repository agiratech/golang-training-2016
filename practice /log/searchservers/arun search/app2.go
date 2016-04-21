package main
import (
	"log"
	"net/http"
	"fmt"

)

func main() {
	http.Handle("/",http.FileServer(http.Dir("stat")))
	http.HandleFunc("http://localhost:3001/arun",handler)
	log.Fatal(http.ListenAndServe(":4001",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "hello, ")
}