package main

import (
    "log"

    "github.com/gin-gonic/gin"
   // "github.com/jinzhu/gorm"
  //  _ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }

func SomeHandler(db *sql.DB) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        // Your handler code goes in here - e.g.
        rows, err := db.Query("select * from student")

        c.String(200, results)
    }

    return gin.HandlerFunc(gn)
}

func main() {
    db, err := sql.Open("postgres", "user=postgres password='root' dbname=aaron sslmode=disable")
    defer db.Close()
    checkErr(err)
    router := gin.Default()
    router.GET("/test", SomeHandler(db))
    router.Run(":8071")
}