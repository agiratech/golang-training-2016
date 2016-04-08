// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	var buf bytes.Buffer
// 	logger := log.New(&buf, "logger:-", log.Lshortfile)
// 	logger.Print("Hello, log file!")

// 	fmt.Print(&buf)
// }
package main

import (
    "github.com/scale-it/go-log"
    "os"
)

func main() {
    Log := log.NewStd(os.Stderr, log.Levels.Debug, log.Ldate|log.Lmicroseconds, true)
    Log.Log(10, "non parametrized message") // level, message
    param := 1.23
    Log.Debugf("some message, %d", param)
    Log.Infof("some message, %d", param)
    Log.Warningf("some message, %d", param)
    Log.Errorf("some message, %d", param)
    Log.Criticalf("some message, %d", param)
    // Log.Fatalf("some message, %d", param) // -- this will break process
}