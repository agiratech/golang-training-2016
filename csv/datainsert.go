package main
 
import (
  "os"
  "log"
  "time"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)
 
func main() {
  // create the statement string
  var sStmt string = "insert into sampletab1 (num,name) values ($1, $2)"
  var entries int = 1000
 
  // lazily open db (doesn't truly open until first request)
  db, err := sql.Open("postgres","user=postgres password='root' dbname=sai sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
 
  fmt.Printf("StartTime: %v\n", time.Now())
   t1 := time.Now()
  for i := 0; i < entries; i++ {
 
    stmt, err := db.Prepare(sStmt)
    if err != nil {
      log.Fatal(err)
    }
 
    res, err := stmt.Exec(1, time.Now())
    if err != nil || res == nil {
      log.Fatal(err)
    }
 
    // close statement
    stmt.Close()
  }
 
  // close db
  db.Close()
   t2:= time.Now()
  fmt.Printf("StopTime: %v\n", time.Now())
  d := t2.Sub(t1)
  fmt.Println("Duration : ",d)
  os.Exit(0)
}