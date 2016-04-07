package main

import (
    "log"
    "net/http"
)

type LogRecord struct {
    http.ResponseWriter
    status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
    return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
    r.status = status
    r.ResponseWriter.WriteHeader(status)
}

func WrapHandler(f http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        record := &LogRecord{
            ResponseWriter: w,
        }

        f.ServeHTTP(record, r)

        log.Println("Bad Request ", record.status)

        if record.status == http.StatusBadRequest {
            log.Println("Bad Request ", r)
        }
    }
}

func main() {
    router := http.NewServeMux()

    s := &http.Server{
        Addr:    ":8080",
        Handler: WrapHandler(router),
    }
    log.Fatal(s.ListenAndServe())
}