package main
 
// schema we can use along with some select statements
// create table test ( gopher_id int, created timestamp );
// select * from test order by created asc limit 1;
// select * from test order by created desc limit 1;
// select count(created) from test;
 
import (
  "log"
  "time"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)
 
const (
  gophers = 10
  entries = 10000
)
 
func main() {
 
  // create string to pass
  var sStmt string = "insert into test (gopher_id, created) values ($1, $2)"
 
  // run the insert function using 10 go routines
  for i := 0; i &lt; gophers; i++ {
    // spin up a gopher
    go gopher(i, sStmt)
  }
 
  // this is a simple way to keep a program open
  // the go program will close when a key is pressed
  var input string
  fmt.Scanln(&amp;input)
}
 
func gopher(gopher_id int, sStmt string) {
 
  // lazily open db (doesn't truly open until first request)
  db, err := sql.Open("postgres","host=localhost dbname=testdb sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
 
  fmt.Printf("Gopher Id: %v || StartTime: %v\n",gopher_id, time.Now())
 
  for i := 0; i &lt; entries; i++ {
 
    stmt, err := db.Prepare(sStmt)
    if err != nil {
      log.Fatal(err)
    }
 
    res, err := stmt.Exec(gopher_id, time.Now())
    if err != nil || res == nil {
      log.Fatal(err)
    }
 
    stmt.Close()
 
  }
  // close db
  db.Close()
 
  fmt.Printf("Gopher Id: %v || StopTime: %v\n",gopher_id, time.Now())
 
}