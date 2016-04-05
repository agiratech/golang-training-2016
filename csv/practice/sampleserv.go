package main 
import (
//  "log"
  "net/http"
)


   func main() {
  http.Handle("/", http.FileServer(http.Dir("./static")))
 
  http.ListenAndServe(":4005", nil)
}