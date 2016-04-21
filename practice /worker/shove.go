package main
 
import (
  //"os"
  "log"
  "time"
  "fmt"
  "database/sql"
  _ "github.com/lib/pq"
)
 
type stu struct {
 
 Id int
 Name string
 Dep string

}

func main() {
  var s stu
  fmt.Println("Enter the values")
  fmt.Scanf("%d",&s.Id)
  fmt.Scanf("%s",&s.Name)
  fmt.Scanf("%s",&s.Dep)
  // create the statement string
  var sStmt string = "insert into student values ($1, $2, $3)"
  var entries int = 100000
  // lazily open db (doesn't truly open until first request)
  db, err := sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
 
 fmt.Printf("StartTime: %v\n", time.Now())
 
 for i:=0; i<entries; i++{

  stmt, err := db.Prepare(sStmt)
  if err != nil {
    log.Fatal(err)
  }
 
  res, err := stmt.Exec(&s.Id,&s.Name,&s.Dep)
  if err != nil || res == nil {
    log.Fatal(err)
  }
 
  // close statement
  stmt.Close()
 }
  // close db
  db.Close()
 
  fmt.Printf("StopTime: %v\n", time.Now())
}