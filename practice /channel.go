package main

import (
    "log"
    "time"
"fmt"
)

var ch = make(chan bool, 1)
var a,b int = 2,3  

func main() {
    log.Println("start")

    //simulates button presses
    go w1(ch)
    go w2(ch)
    go w3(ch)
    go w4(ch)

    //background worker processing requests every second
    for j := 0; j < 10; j++ {
        time.Sleep(time.Second)
        for i := 0; i < len(ch); i++ {
            <-ch
        }
    }
}

//every worker sends a signal to the channel
//and runs the logic within the goroutine
//which after some processing time, gets completed
func w1(ch chan bool) {
    ch <- true
    log.Println("w1 sent")
    go func() {
        time.Sleep(2 * time.Second)
        c:= a + b
        fmt.Println(c)
        log.Println("w1 done")
    }()
}

func w2(ch chan bool) {
    ch <- true
    log.Println("w2 sent")
    go func() {
        time.Sleep(2 * time.Second)
        c:= a + b
        fmt.Println(c)
        log.Println("w2 done")
    }()

}

func w3(ch chan bool) {
    ch <- true
    log.Println("w3 sent")
    go func() {
        time.Sleep(2 * time.Second)
        c:= a + b
        fmt.Println(c)
        log.Println("w3 done")
    }()
}

func w4(ch chan bool) {
    ch <- true
    log.Println("w4 sent")
    go func() {
        time.Sleep(2 * time.Second)
        c:= a + b
        fmt.Println(c)
        log.Println("w4 done")
    }()

}
