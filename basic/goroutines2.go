
package main	

import "fmt"

func f(n int) {
  for i := 0; i < 4; i++ {
    fmt.Println(n, ":", i)
  }
}

func main() {
  for i := 0; i < 4; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}