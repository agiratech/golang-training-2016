package arunpackage


import (
//  "os"
//  "encoding/csv"
  "log"
  "fmt"
  "database/sql"
  "encoding/json"
  _ "github.com/lib/pq"
)

type info struct{

Id int
Name string
Dep string

}
const (

db_user = "postgres"
db_password = "root"
db_dbname = "aaron"

)


//return i

func jesse(e info){

  
}


 
func Db_con(j int) string {

 dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",db_user, db_password, db_dbname )
  db, err := sql.Open("postgres",dbinfo)
  if err != nil {
    log.Fatal(err)
  }

//var id int  
//fmt.Println("Enter the value")
//fmt.Scanf("%d",&id)
   
   var i info
  stmt, err := db.Query("select * from student where id = $1",j)
  if err != nil {
    log.Fatal(err)
  }
  for stmt.Next() {

    
    err := stmt.Scan(&i.Id,&i.Name,&i.Dep)
    if err != nil {
    log.Fatal(err)
  }
  fmt.Println(i.Id,i.Name,i.Dep)

  }

  stmt.Close()
    js,_:=json.Marshal(i)
//fmt.Println(string(js))
  return string(js)
    //jesse(i)
}