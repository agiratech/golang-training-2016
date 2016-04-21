package main

  import (
          "net/http"
          "strings"
  )

  func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
          w.Write([]byte("Hello, World!"))
  }

  func ReplyName(w http.ResponseWriter, r *http.Request) {
          w.Header().Set("Access-Control-Allow-Origin", "*")     //<---------- here! 

          URISegments := strings.Split(r.URL.Path, "/")
          w.Write([]byte(URISegments[1]))
  }

  func main() {
          // http.Handler
          mux := http.NewServeMux()
          mux.HandleFunc("/", SayHelloWorld)
          mux.HandleFunc("/replyname", ReplyName)

          http.ListenAndServe(":8080", mux)
  }