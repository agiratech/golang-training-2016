package main

import (
    "fmt"
    "./workerpackage"
)

func main() {
    worker.WorkerFunc = HandlerFunc

    var lines = []string{"foo", "bar", "car"}

    worker.ProcessStringArray(lines)
}

func HandlerFunc(s string) {
   fmt.Printf("Received: %s\n", s)
}