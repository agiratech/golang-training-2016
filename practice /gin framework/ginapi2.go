package main

import (
     "github.com/gin-gonic/gin"
     _"github.com/lib/pq"
     _"database/sql"
     "github.com/jmoiron/sqlx"
     "log"
     "fmt"
)

type info struct {
Id int         `db:"id" json:"id"`
Name string    `db:"name" json:"name"`
Dep string     `db:"dep" json:"dep"`
}

var db *sqlx.DB

func main() {

   router := gin.Default()
   router.GET("/",search)
   router.Run(":8888")
}

func search(c *gin.Context) {

   var i info
   var id int  
   db,err:= sqlx.Connect("postgres", "user=postgres password='root' dbname=aaron sslmode=disable")
   if err!=nil{
    log.Fatal(err)
   }
   fmt.Println("Enter the value")
   fmt.Scanf("%d",&id)
   
   stmt, err := db.Query("select * from student where id = $1",id)
   if err != nil {
    log.Fatal(err)
    }
   for stmt.Next() { 
    err:= stmt.Scan(&i.Id,&i.Name,&i.Dep)
    if err != nil {
    log.Fatal(err)
  }
   fmt.Println(i.Id,i.Name,i.Dep)
   c.JSON(200,i)
 }
}