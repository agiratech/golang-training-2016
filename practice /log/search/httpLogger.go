package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

var logFile *os.File

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(logFile, "%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}

func main() {
    var err error
    logFile, err = os.Create("logfile.txt")
    if err != nil {
        log.Fatal("Log file create:", err)
        return
    }
    defer logFile.Close()
}