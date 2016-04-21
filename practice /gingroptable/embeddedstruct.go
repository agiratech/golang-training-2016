package main

import ("database/sql"
        "github.com/gin-gonic/gin"
        _"github.com/lib/pq"
        "gopkg.in/gorp.v1"
        "fmt"
        "log"
        "strconv")

type Names struct {

	FirstName string  `db:"firstname" json:"firstname"` 
	LastName string   `db:"lastname" json:"lastname"`
}

type Emp struct {

	Id int64   `db:"id" json:"id"`
	Names
}


var i Emp
var dbmap = initDb()


func initDb() *gorp.DbMap {
  
  db,err := sql.Open("postgres","user=postgres password='root' dbname=aaron sslmode=disable")
  if err!=nil {
  	log.Fatal(err)
  }
  dbmap := &gorp.DbMap{Db: db, Dialect:gorp.PostgresDialect{}}
  return dbmap
}

func Createtable(c *gin.Context) {

	dbmap.AddTableWithName(Emp{},"Emp").SetKeys(true,"Id")
    err := dbmap.CreateTablesIfNotExists()
	if err!=nil {
  	log.Fatal(err)
   }	
}

func Insert(c *gin.Context) {
   
   var i Emp
   fmt.Println("Enter the details")
   fmt.Scanf("%d",&i.Id)
   fmt.Scanf("%s",&i.Names.FirstName)
   fmt.Scanf("%s",&i.Names.LastName) 
   //es := &Emp{3, Names{FirstName: "Alice", LastName: "Smith"}}
   //log.Print(es)
   //log.Print(dbmap)
   _,err := dbmap.Exec("insert into emp values($1,$2,$3)",i.Id,i.Names.FirstName,i.Names.LastName)
   if err!=nil {
  	log.Fatal(err)
   }	
 }

 func Selectid(c *gin.Context) {
   
   ide:= c.Param("id")
   zod,_:=strconv.Atoi(ide)
   var emp Emp
   err := dbmap.SelectOne(&emp,"select * from emp where id=$1",zod)
   if err == nil {
   	c.JSON(200,emp)
   } else {
   	c.JSON(404,gin.H{"error":"nothing in table"})
   }
}

func Updateid(c *gin.Context) {
   
   var i Emp
   fmt.Println("Enter the details")
   fmt.Scanf("%d",&i.Id)
   fmt.Scanf("%s",&i.Names.FirstName)
   fmt.Scanf("%s",&i.Names.LastName) 
   r,err:= dbmap.Exec("update emp set firstname=$1,lastname=$2 where id = $3",i.Names.FirstName,i.Names.LastName,i.Id)
   log.Println(r)
    if err != nil {
   	 	log.Fatal(err)
   	 }
   }


func main() {

	dbmap = initDb()
	router := gin.Default()
	router.GET("/",Createtable)
	router.GET("/user/:id",Selectid)
	router.POST("/user",Insert)
	router.PUT("/user",Updateid)
	router.Run(":2221")
}