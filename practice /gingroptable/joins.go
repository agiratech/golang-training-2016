package main

import ("database/sql"
        "github.com/gin-gonic/gin"
        _"github.com/lib/pq"
        "gopkg.in/gorp.v1"
        "fmt"
        "log"
        "strconv"
        )

type Stud struct {
  
  Sid int       `db:"id" json:"id"`
	Sname string  `db:"sname" json:"sname"` 
	Sdep string   `db:"sdep" json:"sdep"`
}

type Teacher struct {

	Tid int       `db:"tid" json:"tid"`
  Tname string  `db:"tname" json:"tname"` 
  Tdep string   `db:"tdep" json:"tdep"`
  Sid int       `db:"sid" json:"sid"`
	
}

type list struct {
  
  Sid int       `db:"id" json:"id"`
  Sname string  `db:"sname" json:"sname"` 
  Sdep string   `db:"sdep" json:"sdep"`
  Tid int       `db:"tid" json:"tid"`
  Tname string  `db:"tname" json:"tname"` 
  Tdep string   `db:"tdep" json:"tdep"`
  Id int       `db:"sid" json:"sid"`
  
}


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

	dbmap.AddTableWithName(Stud{},"Stud").SetKeys(true,"Sid")
  dbmap.AddTableWithName(Teacher{},"Teacher").SetKeys(true,"Tid")
    err := dbmap.CreateTablesIfNotExists()
	if err!=nil {
  	log.Fatal(err)
   }	
}

func Insert(c *gin.Context) {
   
   var i Stud
   fmt.Println("Enter the Student details")
   fmt.Scanf("%d",&i.Sid)
   fmt.Scanf("%s",&i.Sname)
   fmt.Scanf("%s",&i.Sdep) 
   _,err := dbmap.Exec("insert into Stud values($1,$2,$3)",i.Sid,i.Sname,i.Sdep)
   if err!=nil {
    log.Fatal(err)
   } 
   var j Teacher
   fmt.Println("Enter the Teacher details")
   fmt.Scanf("%d",&j.Tid)
   fmt.Scanf("%s",&j.Tname)
   fmt.Scanf("%s",&j.Tdep)
   fmt.Scanf("%d",&j.Sid)
   _,err= dbmap.Exec("insert into Teacher values($1,$2,$3,$4)",j.Tid,j.Tname,j.Tdep,j.Sid)
   if err!=nil {
    log.Fatal(err)
   }  
 }

 func Selectid(c *gin.Context) {
   
   ide:= c.Param("id")
   zod,_:=strconv.Atoi(ide)
  // var s Stud
  // var t Teacher
   var l []list
   _,err:= dbmap.Select(&l,"select s.id,s.sname,s.sdep,t.tid,t.tname,t.tdep,t.sid from stud s,Teacher t where id=$1",zod)
   log.Println(err)
   if err == nil {
   	c.JSON(200,l)
   } else {
   	c.JSON(404,gin.H{"error":"nothing in table"})
   }
}


func main() {

	//dbmap = initDb()
	router := gin.Default()
	router.GET("/",Createtable)
	router.POST("/user",Insert)
  router.GET("/user/:id",Selectid)
	router.Run(":3333")
}