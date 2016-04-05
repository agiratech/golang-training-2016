package main

import (
    "bufio"
    "encoding/csv"
    "os"
    "fmt"
    "io"
)

func main() {

  f, _ := os.Open("/home/newuser/Desktop/myfile_sample.csv")
   
   r := csv.NewReader(bufio.NewReader(f))
    for {
  record, err := r.Read()

  if err == io.EOF {
      break
  }

  fmt.Println(record)
  /* fmt.Println(len(record))
  for value := range record {
      fmt.Printf("  %v\n", record[value])
  } */
    }
}