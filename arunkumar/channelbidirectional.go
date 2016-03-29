package main

import ("fmt"
        "time"
)

func main() {
  c1 := make(chan string,2)
  c2 := make(chan string,2)

 go func() {
    for {
      c2 <- "from 2"
      c2 <- "from 6"

      time.Sleep(1000000000)
    }
  }()

  go func() {
    for {
      c1 <- "from 1"
      c1 <- "from 5"

      time.Sleep(1000000000)
    }
  }()

 

  go func() {
    for {
      select {
      case msg1 := <- c1:
        fmt.Println(msg1)
      case msg2 := <- c2:
        fmt.Println(msg2)
         case msg3 := <- c1:
        fmt.Println(msg3)
      case msg4 := <- c2:
        fmt.Println(msg4)
        case <- time.After(time.Second):
     fmt.Println("timeout")
 // default:
  //fmt.Println("nothing ready")
      }
    }
  }()

  var input string
  fmt.Scanln(&input)
}