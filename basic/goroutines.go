package main

import (
	"fmt"
	"sync"
	// "time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s ---->%d \n",s,i)
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	go say("world")
	say("hello")
	wg.Wait()
}
