package main

import ("database/sql"
        _"github.com/gin-gonic/gin"
        _"github.com/lib/pq"
        "gopkg.in/gorp.v1"
        "log")


type User struct{

	Id int
	Name string
	Dep string
}

func checkerr(err error, msg string){
 
	if err!=nil {
	log.Fatal(err,msg)
	}
}

func initDb() *gorp.DbMap {
	
	 db, err := sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
     checkerr(err,"sql.Open failed")    
	 dbmap := &gorp.DbMap{Db: db, Dialect:gorp.PostgresDialect{}}
     dbmap.AddTableWithName(User{}, "arun").SetKeys(true, "Id")
     err = dbmap.CreateTablesIfNotExists()
     checkerr(err,"Create table failed")

     return dbmap 

}

func main() {

	dbmap := initDb()
	defer dbmap.Db.Close()
}


